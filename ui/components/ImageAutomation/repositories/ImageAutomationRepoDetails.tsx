import * as React from "react";
import styled from "styled-components";
import { useFeatureFlags } from "../../../hooks/featureflags";
import { useGetObject } from "../../../hooks/objects";
import { Kind } from "../../../lib/api/core/types.pb";
import { ImageRepository } from "../../../lib/objects";
import { V2Routes } from "../../../lib/types";
import Button from "../../Button";
import Interval from "../../Interval";
import Link from "../../Link";
import Page from "../../Page";
import ImageAutomationDetails from "../ImageAutomationDetails";

type Props = {
  className?: string;
  name: string;
  namespace: string;
  clusterName: string;
};

function ImageAutomationRepoDetails({
  className,
  name,
  namespace,
  clusterName,
}: Props) {
  const { data, isLoading, error } = useGetObject<ImageRepository>(
    name,
    namespace,
    Kind.ImageRepository,
    clusterName,
    {
      refetchInterval: 5000,
    }
  );
  const { isFlagEnabled } = useFeatureFlags();
  const rootPath = V2Routes.ImageAutomationRepositoryDetails;
  return (
    <Page
      error={error}
      loading={isLoading}
      className={className}
      path={[
        { label: "Image Repositories", url: V2Routes.ImageRepositories },
        { label: name },
      ]}
    >
      {!!data && (
        <ImageAutomationDetails
          data={data}
          kind={Kind.ImageRepository}
          infoFields={[
            ["Kind", Kind.ImageRepository],
            ["Namespace", data.namespace],
            isFlagEnabled("WEAVE_GITOPS_FEATURE_CLUSTER") && [
              "Cluster",
              clusterName,
            ],
            [
              "Image",
              <Link newTab={true} to={data.obj?.spec?.image}>
                {data.obj?.spec?.image}
              </Link>,
            ],
            ["Interval", <Interval interval={data.interval} />],
            ["Tag Count", data.tagCount],
          ]}
          rootPath={rootPath}
        >
          <Button>
            <Link
              to={`/image_automation/policies?filters=imageRepositoryRef: ${name}_`}
            >
              Go To Image Policy
            </Link>
          </Button>
        </ImageAutomationDetails>
      )}
    </Page>
  );
}

export default styled(ImageAutomationRepoDetails)``;
