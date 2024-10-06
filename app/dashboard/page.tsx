import { EstablishmentList } from "@/components/establishments/EstablishmentList";

export default function Dashboard() {
  return (
    <div>
      <h1 className="text-3xl font-bold mb-6">Bars & Restaurants in Ginza</h1>
      <EstablishmentList />
    </div>
  );
}