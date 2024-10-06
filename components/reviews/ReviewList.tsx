import { Review } from "@/lib/types";

interface ReviewListProps {
  reviews: Review[];
}

export function ReviewList({ reviews }: ReviewListProps) {
  return (
    <div className="space-y-4">
      {reviews.map((review) => (
        <div key={review.id} className="border p-4 rounded">
          <p>Rating: {review.rating}/5</p>
          <p>{review.content}</p>
        </div>
      ))}
    </div>
  );
}