import { refreshAccessToken, getFullUrl } from "redux-data-connect";

export const isDev = process.env.NODE_ENV === "development";

export const GANACHE_NETWORK_ID = 5777;
export const KOVAN_NETWORK_ID = 42;
export const CHAIN_ID = KOVAN_NETWORK_ID;

export const APP_VERSION = "1";

export const CONTRACT_ADDRESS = "0x1c56346cd2a2bf3202f771f50d3d14a367b48070";
export const CONTRACT_SALT = "";

export const BACKEND_ADDR = isDev
  ? "http://localhost:8080"
  : "https://juicer.finance";

export const GA_TRACKER = "UA-178882519-1";
export const FB_PIXEL = "297559898107412"; //"1039111996558849";
export const SENTRY_DSN =
  "https://cd19416ad99349d0bc8df4b50b374d4e@sentry.io/3026714";

export const AUTH_API = {
  GoogleRedirectEndpoint: getFullUrl("/auth/login/google/code/", {
    host: BACKEND_ADDR,
  }),
  GoogleDoneEndpoint: getFullUrl("/auth/login/google/done/", {
    host: BACKEND_ADDR,
  }),
  RefreshTokenEndpoint: getFullUrl("/auth/token/refresh/", {
    host: BACKEND_ADDR,
  }),
};
