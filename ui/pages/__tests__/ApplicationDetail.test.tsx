import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import "jest-styled-components";
import * as React from "react";
import { act } from "react-dom/test-utils";
import {
  ListCommitsResponse,
  SyncApplicationResponse,
} from "../../lib/api/applications/applications.pb";
import { withContext, withTheme } from "../../lib/test-utils";
import ApplicationDetail from "../ApplicationDetail";

describe("ApplicationDetail", () => {
  describe("Sync App Button", () => {
    const apiMock = {
      GetApplication: () => ({
        application: {
          name: "pod-info",
          namespace: "wego-systems",
        },
      }),
      ListCommits: (): ListCommitsResponse => ({
        commits: [
          {
            hash: "123abc",
            author: "Example User",
            date: "2021-09-10T23:45:09Z",
          },
        ],
      }),
      SyncApplication: (): SyncApplicationResponse => {
        return {
          success: true,
        };
      },
    };

    it("should exist on page", async () => {
      render(
        withTheme(
          withContext(
            <ApplicationDetail name="pod-info" />,
            "/application_detail",
            apiMock
          )
        )
      );
      expect(await screen.findByText("Sync App")).toBeTruthy();
    });
    it("should call a sync request", async () => {
      const promise = Promise.resolve();

      apiMock.SyncApplication = jest.fn();

      await act(async () => {
        render(
          withTheme(
            withContext(
              <ApplicationDetail name="pod-info" />,
              "/application_detail",
              apiMock
            )
          )
        );
      });

      const button = await (
        await screen.findByText("Sync App")
      ).closest("button");

      fireEvent(button, new MouseEvent("click", { bubbles: true }));

      expect(apiMock.SyncApplication).toHaveBeenCalledWith({
        name: "pod-info",
        namespace: "wego-systems",
      });
      await waitFor(() => promise);
    });
    // it("should notify user on success", () => "");
    // it("should notify user on failure", () => "");
  });
});
