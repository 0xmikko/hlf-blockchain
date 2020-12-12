export const isDev = process.env.NODE_ENV === "development";

// Backend address of node on Network #1
export const BACKEND_ADDR_1 = isDev
  ? "http://localhost:8080"
  : "https://juicer.finance";

// Backend address of node on Network #2
export const BACKEND_ADDR_2 = isDev
    ? "http://localhost:8080"
    : "https://juicer.finance";
