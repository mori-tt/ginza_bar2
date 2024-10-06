import Link from 'next/link';

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <h1 className="text-4xl font-bold">Welcome to Ginza Guide</h1>
      <p className="mt-4 text-xl">Discover the best bars and restaurants in Ginza, Tokyo</p>
      <Link href="/dashboard" className="mt-8 px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600">
        Go to Dashboard
      </Link>
    </main>
  );
}