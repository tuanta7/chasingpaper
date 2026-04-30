import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/subscriptions")({
  component: SubscriptionsPage,
});

function SubscriptionsPage() {
  return (
    <div className="flex flex-col items-center justify-center p-10 text-center">
      <h2 className="text-2xl font-bold">Subscriptions</h2>
    </div>
  );
}
