import _ from "lodash";
import * as React from "react";
import { useRouteMatch } from "react-router-dom";
import styled from "styled-components";
import { useListAutomations } from "../hooks/automations";
import { Kind } from "../lib/api/core/types.pb";
import { HelmRelease, Source } from "../lib/objects";
import { getSourceRefForAutomation } from "../lib/utils";
import AutomationsTable from "./AutomationsTable";
import EventsTable from "./EventsTable";
import Flex from "./Flex";
import InfoList, { InfoField } from "./InfoList";
import LoadingPage from "./LoadingPage";
import Metadata from "./Metadata";
import PageStatus from "./PageStatus";
import SubRouterTabs, { RouterTab } from "./SubRouterTabs";
import SyncActions from "./SyncActions";
import YamlView from "./YamlView";

type Props = {
  className?: string;
  type: Kind;
  children?: JSX.Element;
  source: Source;
  info: InfoField[];
  customActions?: JSX.Element[];
};

function SourceDetail({ className, source, info, type, customActions }: Props) {
  const { name, namespace, clusterName, suspended } = source;

  const { data: automations, isLoading: automationsLoading } =
    useListAutomations();
  const { path } = useRouteMatch();

  if (automationsLoading) {
    return <LoadingPage />;
  }

  const isNameRelevant = (expectedName) => {
    return expectedName == source.name;
  };

  const isRelevant = (expectedType, expectedName) => {
    return expectedType == source.type && isNameRelevant(expectedName);
  };

  const relevantAutomations = _.filter(automations?.result, (a) => {
    if (!source) {
      return false;
    }
    if (a.clusterName != source.clusterName) {
      return false;
    }

    if (type === Kind.HelmChart) {
      return isNameRelevant((a as HelmRelease)?.helmChart?.name);
    }

    const sourceRef = getSourceRefForAutomation(a);

    return isRelevant(sourceRef?.kind, sourceRef?.name);
  });

  return (
    <Flex wide tall column className={className} gap="32">
      <Flex column gap="8">
        <PageStatus
          conditions={source.conditions}
          suspended={source.suspended}
        />
        <SyncActions
          name={name}
          namespace={namespace}
          clusterName={clusterName}
          kind={type}
          suspended={suspended}
          hideDropdown
          customActions={customActions}
        />
      </Flex>

      <SubRouterTabs rootPath={`${path}/details`}>
        <RouterTab name="Details" path={`${path}/details`}>
          <>
            <InfoList items={info} />
            <Metadata metadata={source.metadata} labels={source.labels} />
            <AutomationsTable automations={relevantAutomations} hideSource />
          </>
        </RouterTab>
        <RouterTab name="Events" path={`${path}/events`}>
          <EventsTable
            namespace={source.namespace}
            involvedObject={{
              kind: source.type,
              name: source.name,
              namespace: source.namespace,
              clusterName: source.clusterName,
            }}
          />
        </RouterTab>
        <RouterTab name="yaml" path={`${path}/yaml`}>
          <YamlView
            yaml={source.yaml}
            object={{
              kind: source.type,
              name: source.name,
              namespace: source.namespace,
            }}
          />
        </RouterTab>
      </SubRouterTabs>
    </Flex>
  );
}

export default styled(SourceDetail).attrs({ className: SourceDetail.name })`
  ${SubRouterTabs} {
    margin-top: ${(props) => props.theme.spacing.medium};
  }
`;
