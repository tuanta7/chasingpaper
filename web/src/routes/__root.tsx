import { useState } from "react";
import { Button, Layout, Menu } from "antd";
import {
  Bars3Icon,
  ChevronLeftIcon,
  HomeIcon,
  QueueListIcon,
  CreditCardIcon,
} from "@heroicons/react/24/outline";
import {
  Link,
  Outlet,
  createRootRoute,
  useRouterState,
} from "@tanstack/react-router";
import { Toaster } from "sonner";

const { Sider, Content } = Layout;

export const Route = createRootRoute({
  component: RootLayout,
});

function RootLayout() {
  const [selectedKey, setSelectedKey] = useState("/");

  const [open, setOpen] = useState(true);

  return (
    <Layout className="h-screen overflow-hidden bg-cream-100 font-sans">
      <Sider
        theme="light"
        width={240}
        breakpoint="lg"
        collapsedWidth={72}
        collapsible
        collapsed={!open}
        onBreakpoint={(broken) => {
          setOpen(!broken); // Automatically collapse on smaller screens and expand on larger screens
        }}
        trigger={null}
        className="border-r border-cream-200"
      >
        <div
          className={`flex items-center py-6 ${open ? "pl-5 pr-1 justify-between" : "justify-center"}`}
        >
          {open && (
            <div className="mr-2 truncate text-xl font-extrabold tracking-tight text-navy-800">
              Chasing Control
            </div>
          )}
          <Button
            type="text"
            onClick={() => setOpen((value) => !value)}
            className="flex shrink-0 items-center justify-center px-2! text-navy-700 hover:bg-cream-200! hover:text-navy-900!"
          >
            {open ? (
              <ChevronLeftIcon className="h-4 w-4" />
            ) : (
              <Bars3Icon className="h-5 w-5" />
            )}
          </Button>
        </div>
        <Menu
          mode="inline"
          selectedKeys={[selectedKey]}
          className="border-none px-3 font-semibold"
          items={[
            {
              key: "/",
              icon: <HomeIcon className="h-5" />,
              label: (
                <Link to="/" onClick={() => setSelectedKey("/")}>
                  Home
                </Link>
              ),
            },
            {
              key: "/plans",
              icon: <QueueListIcon className="h-5" />,
              label: (
                <Link to="/plans" onClick={() => setSelectedKey("/plans")}>
                  Plans
                </Link>
              ),
            },
            {
              key: "/subscriptions",
              icon: <CreditCardIcon className="h-5" />,
              label: (
                <Link
                  to="/subscriptions"
                  onClick={() => setSelectedKey("/subscriptions")}
                >
                  Subscriptions
                </Link>
              ),
            },
          ]}
        />
      </Sider>
      <Layout className="flex-1 overflow-y-auto bg-cream-100">
        <Content className="px-4 py-8 md:px-10 md:py-10">
          <div className="mx-auto max-w-5xl text-navy-800">
            <Outlet />
          </div>
        </Content>
      </Layout>
      <Toaster />
    </Layout>
  );
}
