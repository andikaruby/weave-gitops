import {
  Checkbox,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  withStyles,
} from "@material-ui/core";
import _ from "lodash";
import * as React from "react";
import { useHistory } from "react-router-dom";
import styled from "styled-components";
import { theme } from "..";
import Button from "./Button";
import Flex from "./Flex";
import Icon, { IconType } from "./Icon";
import Spacer from "./Spacer";
import Text from "./Text";

export enum SortType {
  //sort is unused but having number as index zero makes it a falsy value thus not used as a valid sortType for selecting fields for SortableLabel
  sort,
  number,
  string,
  date,
  bool,
}

type Sorter = (k: any) => any;

export type Field = {
  label: string;
  value: string | ((k: any) => string | JSX.Element);
  sortType?: SortType;
  sortValue?: Sorter;
  textSearchable?: boolean;
  maxWidth?: number;
};

/** DataTable Properties  */
export interface Props {
  /** CSS MUI Overrides or other styling. */
  className?: string;
  /** A list of objects with four fields: `label`, which is a string representing the column header, `value`, which can be a string, or a function that extracts the data needed to fill the table cell, `sortType`, which determines the sorting function to be used, and `sortValue`, which customizes your input to the search function */
  fields: Field[];
  /** A list of data that will be iterated through to create the columns described in `fields`. */
  rows?: any[];
  /** index of field to initially sort against. */
  defaultSort?: number;
  /** for passing pagination */
  children?: any;

  selectable?: any;
  selectedRows?: any;
  setSelectedRows?: any;
}

const EmptyRow = styled(TableRow)<{ colSpan: number }>`
  td {
    text-align: center;
  }
`;

const TableButton = styled(Button)`
  &.MuiButton-root {
    margin: 0;
    text-transform: none;
  }
  &.MuiButton-text {
    min-width: 0px;
    .selected {
      color: ${(props) => props.theme.colors.neutral40};
    }
  }
  &.arrow {
    min-width: 0px;
  }
`;

type Row = any;

function defaultSortFunc(sort: Field): Sorter {
  return (a: Row) => {
    return a[sort.value as string];
  };
}

export const sortWithType = (rows: Row[], sort: Field) => {
  const sortFn = sort.sortValue || defaultSortFunc(sort);
  return (rows || []).sort((a: Row, b: Row) => {
    switch (sort.sortType) {
      case SortType.number:
        return sortFn(a) - sortFn(b);

      case SortType.date:
        return Date.parse(sortFn(a)) - Date.parse(sortFn(b));

      case SortType.bool:
        if (sortFn(a) === sortFn(b)) return 0;
        else if (sortFn(a) === false && sortFn(b) === true) return -1;
        else return 1;

      default:
        return (sortFn(a) || "").localeCompare(sortFn(b) || "");
    }
  });
};

/** Form DataTable */
function UnstyledDataTable({
  className,
  fields,
  rows,
  defaultSort = 0,
  children,
  selectable,
  selectedRows,
  setSelectedRows,
}: Props) {
  const [sort, setSort] = React.useState(fields[defaultSort]);
  const [reverseSort, setReverseSort] = React.useState(false);
  const history = useHistory();

  const sorted = sortWithType(rows, sort);

  if (reverseSort) {
    sorted.reverse();
  }

  type labelProps = {
    field: Field;
  };
  function SortableLabel({ field }: labelProps) {
    return (
      <Flex align start>
        <TableButton
          color="inherit"
          variant="text"
          onClick={() => {
            setReverseSort(sort === field ? !reverseSort : false);
            setSort(field);
          }}
        >
          <h2 className={sort.label === field.label ? "selected" : ""}>
            {field.label}
          </h2>
        </TableButton>
        <Spacer padding="xxs" />
        {sort.label === field.label ? (
          <Icon
            type={IconType.ArrowUpwardIcon}
            size="base"
            className={reverseSort ? "upward" : "downward"}
          />
        ) : (
          <div style={{ width: "16px" }} />
        )}
      </Flex>
    );
  }

  const IndividualCheckbox = withStyles({
    root: {
      color: theme.colors.primary,
      "&$checked": {
        color: theme.colors.primary,
      },
      "&$disabled": {
        color: theme.colors.neutral20,
      },
    },
    checked: {},
    disabled: {},
  })(Checkbox);

  const handleSelectAllClick = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.checked) {
      const newSelected = rows.map((row, i) => i) || [];
      setSelectedRows(newSelected);
      return;
    }
    setSelectedRows([]);
  };

  const handleClick = (event: React.MouseEvent<unknown>, name: string) => {
    const selectedIndex = selectedRows.indexOf(name);
    let newSelected: string[] = [];

    console.log(name);

    if (selectedIndex === -1) {
      newSelected = newSelected.concat(selectedRows, name);
    } else if (selectedIndex === 0) {
      newSelected = newSelected.concat(selectedRows.slice(1));
    } else if (selectedIndex === selectedRows.length - 1) {
      newSelected = newSelected.concat(selectedRows.slice(0, -1));
    } else if (selectedIndex > 0) {
      newSelected = newSelected.concat(
        selectedRows.slice(0, selectedIndex),
        selectedRows.slice(selectedIndex + 1)
      );
    }
    setSelectedRows(newSelected);
  };

  const numSelected = (selectedRows || []).length;
  const rowCount = rows.length || 0;
  const isSelected = (name: string) => selectedRows.indexOf(name) !== -1;

  const r = _.map(sorted, (r, i) => {
    const labelId = `enhanced-table-checkbox-${r.name}`;
    const isItemSelected = isSelected(r.name);
    return (
      <TableRow key={i}>
        {/* hide behind selectable attribute */}
        <TableCell padding="checkbox">
          <IndividualCheckbox
            checked={isItemSelected}
            inputProps={{ "aria-labelledby": labelId }}
            onClick={(event: any) => handleClick(event, r.name)}
          />
        </TableCell>
        {_.map(fields, (f) => (
          <TableCell
            style={
              f.maxWidth && {
                maxWidth: f.maxWidth,
              }
            }
            key={f.label}
          >
            <Text>
              {typeof f.value === "function" ? f.value(r) : r[f.value]}
            </Text>
          </TableCell>
        ))}
      </TableRow>
    );
  });

  React.useEffect(() => {
    return history.listen(() => {
      setSelectedRows([]);
    });
  }, [setSelectedRows]);

  return (
    <div className={className}>
      <TableContainer>
        <Table aria-label="simple table">
          <TableHead>
            <TableRow>
              {/* hide behind selectable */}
              <TableCell padding="checkbox">
                <Checkbox
                  indeterminate={numSelected > 0 && numSelected < rowCount}
                  checked={rowCount > 0 && numSelected === rowCount}
                  onChange={handleSelectAllClick}
                  inputProps={{ "aria-label": "select all rows" }}
                  style={{
                    color: theme.colors.primary,
                  }}
                />
              </TableCell>
              {_.map(fields, (f) => (
                <TableCell key={f.label}>
                  <SortableLabel field={f} />
                </TableCell>
              ))}
            </TableRow>
          </TableHead>
          <TableBody>
            {r.length > 0 ? (
              r
            ) : (
              <EmptyRow colSpan={fields.length}>
                <TableCell colSpan={fields.length}>
                  <Flex center align>
                    <Icon
                      color="neutral20"
                      type={IconType.RemoveCircleIcon}
                      size="base"
                    />
                    <Spacer padding="xxs" />
                    <Text color="neutral30">No data</Text>
                  </Flex>
                </TableCell>
              </EmptyRow>
            )}
          </TableBody>
        </Table>
      </TableContainer>
      <Spacer padding="xs" />
      {/* optional pagination component */}
      {children}
    </div>
  );
}

export const DataTable = styled(UnstyledDataTable)`
  width: 100%;
  overflow-x: auto;
  h2 {
    padding: ${(props) => props.theme.spacing.xs};
    font-size: 18px;
    font-weight: 600;
    color: ${(props) => props.theme.colors.neutral30};
    margin: 0px;
    white-space: nowrap;
  }
  .MuiTableRow-root {
    transition: background 0.5s ease-in-out;
  }
  .MuiTableRow-root:not(.MuiTableRow-head):hover {
    background: ${(props) => props.theme.colors.neutral10};
    transition: background 0.5s ease-in-out;
  }
  th {
    padding: 0;
  }
  td {
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }
`;

export default DataTable;
