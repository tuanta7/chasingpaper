import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { QueryClientProvider, QueryClient } from "@tanstack/react-query";
import { RouterProvider } from "@tanstack/react-router";
import { ConfigProvider } from "antd";

import { router } from "./router";
import "./index.css";

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: 1,
      staleTime: 30_000,
    },
  },
});

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <ConfigProvider
      theme={{
        token: {
          colorPrimary: "#1a252f",
          colorBgBase: "#fdfbf7",
          colorTextBase: "#1a252f",
          fontFamily: "'Nunito', sans-serif",
          borderRadius: 16,
        },
        components: {
          Layout: {
            bodyBg: "#fdfbf7",
            headerBg: "#fdfbf7",
            siderBg: "#fdfbf7",
          },
          Menu: {
            itemBg: "transparent",
            itemColor: "#2c3e50",
            itemSelectedBg: "#f4efdf",
            itemSelectedColor: "#0d131a",
          },
        },
      }}
    >
      <QueryClientProvider client={queryClient}>
        <RouterProvider router={router} />
      </QueryClientProvider>
    </ConfigProvider>
  </StrictMode>,
);
