export interface Establishment {
  id: number;
  name: string;
  description: string;
  rating: number;
}

export interface Review {
  id: number;
  establishmentId: number;
  rating: number;
  content: string;
}

export interface User {
  id: number;
  name: string;
  email: string;
  role: "super" | "regular";
  approved: boolean;
}