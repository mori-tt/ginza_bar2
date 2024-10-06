import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Establishment } from "@/lib/types";

interface EstablishmentCardProps {
  establishment: Establishment;
}

export function EstablishmentCard({ establishment }: EstablishmentCardProps) {
  return (
    <Card>
      <CardHeader>
        <CardTitle>{establishment.name}</CardTitle>
      </CardHeader>
      <CardContent>
        <p>{establishment.description}</p>
        <p>Rating: {establishment.rating}/5</p>
      </CardContent>
    </Card>
  );
}