import { createFileRoute } from "@tanstack/react-router";
import { SparklesIcon } from "@heroicons/react/24/solid";

export const Route = createFileRoute("/")({
  component: IndexComponent,
});

function IndexComponent() {
  return (
    <div className="flex min-h-100 flex-col items-center justify-center rounded-3xl border border-cream-200 bg-white p-10 text-center shadow-sm">
      <div className="mb-6 flex h-20 w-20 items-center justify-center rounded-full bg-cream-100 text-navy-700 shadow-inner">
        <SparklesIcon className="h-10 w-10" />
      </div>
      <h2 className="text-3xl font-extrabold text-navy-800">
        Welcome to Chasing Control
      </h2>
      <p className="max-w-md mt-4 text-base font-medium text-navy-700/80">
        This is your dashboard!
      </p>
    </div>
  );
}
