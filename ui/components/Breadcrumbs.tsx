import { Tooltip } from "@material-ui/core";
import React from "react";
import styled from "styled-components";
import Flex from "./Flex";
import Icon, { IconType } from "./Icon";
import Link from "./Link";
import Text from "./Text";

const EllipsesText = styled(Text)<{ maxWidth?: string }>`
  max-width: ${(prop) => prop.maxWidth || "400px"};
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
`;
export interface Breadcrumb {
  label: string;
  url?: string;
}
interface Props {
  path: Breadcrumb[];
}
export const Breadcrumbs = ({ path = [] }: Props) => {
  return (
    <Flex align>
      {path.map(({ label, url }) => {
        return (
          <Flex align key={label}>
            {url ? (
              <>
                <Link
                  data-testid={`link-${label}`}
                  to={url}
                  textProps={{ bold: true, size: "large", color: "black" }}
                >
                  {label}
                </Link>
                <Icon
                  type={IconType.NavigateNextIcon}
                  size="large"
                  color="black"
                />
              </>
            ) : (
              <Tooltip title={label} placement="bottom">
                <EllipsesText
                  size="large"
                  color="black"
                  className="ellipsis"
                  data-testid={`text-${label}`}
                >
                  {label}
                </EllipsesText>
              </Tooltip>
            )}
          </Flex>
        );
      })}
    </Flex>
  );
};

export default styled(Breadcrumbs).attrs({ className: Breadcrumbs.name })``;
