"use client";

import { useState, useEffect } from "react";
import { EstablishmentCard } from "./EstablishmentCard";
import { Establishment } from "@/lib/types";

export function EstablishmentList() {
  const [establishments, setEstablishments] = useState<Establishment[]>([]);

  useEffect(() => {
    // Fetch establishments from the API
    // This is a placeholder and should be replaced with actual API call
    setEstablishments([
      { id: 1, name: "Bar A", description: "A cozy bar", rating: 4.5 },
      { id: 2, name: "Restaurant B", description: "Fine dining", rating: 4.8 },
    ]);
  }, []);

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      {establishments.map((establishment) => (
        <EstablishmentCard key={establishment.id} establishment={establishment} />
      ))}
    </div>
  );
}