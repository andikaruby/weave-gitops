import * as React from "react";
import { useRouteMatch } from "react-router-dom";
import styled from "styled-components";
import { AppContext } from "../contexts/AppContext";
import { useSyncAutomation } from "../hooks/automations";
import {
  AutomationKind,
  HelmRelease,
  SourceRefSourceKind,
} from "../lib/api/core/types.pb";
import Alert from "./Alert";
import EventsTable from "./EventsTable";
import Flex from "./Flex";
import InfoList from "./InfoList";
import Interval from "./Interval";
import PageStatus from "./PageStatus";
import ReconciledObjectsTable from "./ReconciledObjectsTable";
import SourceLink from "./SourceLink";
import SubRouterTabs, { RouterTab } from "./SubRouterTabs";
import SyncButton from "./SyncButton";

type Props = {
  name: string;
  clusterName: string;
  helmRelease?: HelmRelease;
  className?: string;
};

const TabContent = styled.div`
  margin-top: ${(props) => props.theme.spacing.medium};
  width: 100%;
  height: 100%;
`;

function helmChartLink(helmRelease: HelmRelease) {
  if (helmRelease.helmChartName === "") {
    return (
      <SourceLink
        sourceRef={{
          kind: SourceRefSourceKind.HelmChart,
          name: helmRelease?.helmChart.chart,
        }}
      />
    );
  }

  const [ns, name] = helmRelease.helmChartName.split("/");

  return (
    <SourceLink
      sourceRef={{
        kind: SourceRefSourceKind.HelmChart,
        name: name,
        namespace: ns,
      }}
    />
  );
}

function HelmReleaseDetail({ name, helmRelease, className }: Props) {
  const { path } = useRouteMatch();
  const { notifySuccess } = React.useContext(AppContext);
  const sync = useSyncAutomation({
    name: helmRelease?.name,
    namespace: helmRelease?.namespace,
    clusterName: helmRelease?.clusterName,
    kind: AutomationKind.HelmReleaseAutomation,
  });

  const handleSyncClicked = (opts) => {
    sync.mutateAsync(opts).then(() => {
      notifySuccess("Resource synced successfully");
    });
  };

  return (
    <Flex wide tall column className={className}>
      {sync.isError && (
        <Alert
          severity="error"
          message={sync.error.message}
          title="Sync Error"
        />
      )}
      <PageStatus
        conditions={helmRelease?.conditions}
        suspended={helmRelease?.suspended}
      />
      <InfoList
        items={[
          ["Namespace", helmRelease?.namespace],
          ["Source", helmChartLink(helmRelease)],
          ["Chart", helmRelease?.helmChart.chart],
          ["Cluster", helmRelease?.clusterName],
          ["Interval", <Interval interval={helmRelease?.interval} />],
        ]}
      />
      <SyncButton onClick={handleSyncClicked} loading={sync.isLoading} />
      <TabContent>
        <SubRouterTabs rootPath={`${path}/details`}>
          <RouterTab name="Details" path={`${path}/details`}>
            <ReconciledObjectsTable
              kinds={helmRelease?.inventory}
              automationName={helmRelease?.name}
              automationKind={AutomationKind.HelmReleaseAutomation}
              clusterName={helmRelease?.clusterName}
            />
          </RouterTab>
          <RouterTab name="Events" path={`${path}/events`}>
            <EventsTable
              involvedObject={{
                kind: "HelmRelease",
                name,
                namespace: helmRelease?.namespace,
              }}
            />
          </RouterTab>
        </SubRouterTabs>
      </TabContent>
    </Flex>
  );
}

export default styled(HelmReleaseDetail).attrs({
  className: HelmReleaseDetail.name,
})`
  width: 100%;

  ${Alert} {
    margin-bottom: 16px;
  }
`;
