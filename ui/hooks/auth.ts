import { useContext, useState } from "react";
import { AppContext } from "../contexts/AppContext";
import {
  GetGithubAuthStatusResponse,
  GetGithubDeviceCodeResponse,
} from "../lib/api/applications/applications.pb";
import { GrpcErrorCodes } from "../lib/types";

function poller(cb, interval) {
  if (process.env.NODE_ENV === "test") {
    // Stay synchronous in tests
    return cb();
  }

  return setInterval(cb, interval);
}

export function isUnauthenticated(code: GrpcErrorCodes) {
  return code === GrpcErrorCodes.Unauthenticated;
}

export default function useAuth() {
  const [loading, setLoading] = useState(true);
  const { applicationsClient, getProviderToken, storeProviderToken } =
    useContext(AppContext);

  const getGithubDeviceCode = () => {
    setLoading(true);
    return applicationsClient
      .GetGithubDeviceCode({})
      .finally(() => setLoading(false));
  };

  const getGithubAuthStatus = (codeRes: GetGithubDeviceCodeResponse) => {
    let poll;
    return {
      cancel: () => clearInterval(poll),
      promise: new Promise<GetGithubAuthStatusResponse>((accept, reject) => {
        poll = poller(() => {
          applicationsClient
            .GetGithubAuthStatus(codeRes)
            .then((res) => {
              clearInterval(poll);
              accept(res);
            })
            .catch(({ code, message }) => {
              // Unauthenticated means we can keep polling.
              //  On anything else, stop polling and report.
              if (!isUnauthenticated(code)) {
                clearInterval(poll);
                reject({ message });
              }
            });
        }, (codeRes.interval + 1) * 1000);
      }),
    };
  };

  return {
    loading,
    getGithubDeviceCode,
    getGithubAuthStatus,
    getProviderToken,
    storeProviderToken,
  };
}
