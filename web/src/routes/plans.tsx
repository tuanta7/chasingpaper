import { createFileRoute } from "@tanstack/react-router";
import { usePlansQuery } from "../data/plans/list";

export const Route = createFileRoute("/plans")({
  component: PlansPage,
});

function PlansPage() {
  const { data: plans = [], isLoading, isError, error } = usePlansQuery();

  if (isLoading) {
    return (
      <section className="space-y-4">
        <p className="text-navy-700 font-medium">Baking plans...</p>
      </section>
    );
  }

  if (isError) {
    return (
      <section className="space-y-4">
        <div className="rounded-2xl bg-red-50 p-4 border border-red-100 text-red-600">
          Oops! {error.message}
        </div>
      </section>
    );
  }

  return (
    <section className="space-y-6">
      {plans.length === 0 ? (
        <div className="rounded-3xl border-2 border-dashed border-cream-300 p-10 text-center">
          <p className="text-lg font-bold text-navy-700">No plans found.</p>
          <p className="text-navy-700/60 mt-1">Shall we create the first one?</p>
        </div>
      ) : (
        <ul className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
          {plans.map((plan, index) => (
            <li
              key={String(plan.id ?? index)}
              className="group relative rounded-3xl border border-cream-200 bg-white p-6 shadow-sm transition-all hover:shadow-md hover:-translate-y-1"
            >
              <h3 className="font-bold text-lg text-navy-800">
                {String(plan.title ?? plan.name ?? `Plan ${index + 1}`)}
              </h3>
              <p className="mt-2 text-sm font-medium text-navy-700/70">
                Some little details here...
              </p>
            </li>
          ))}
        </ul>
      )}
    </section>
  );
}
