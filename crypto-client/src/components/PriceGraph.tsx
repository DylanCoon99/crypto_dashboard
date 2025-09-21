// crypto-dashboard-client/components/PriceGraph.tsx
'use client';

import { useQuery } from '@tanstack/react-query';
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer } from 'recharts';
import { fetchPriceHistory } from '@/lib/api';

export default function PriceGraph({ coinId }: { coinId: string }) {
  const { data, isLoading, error } = useQuery({
    queryKey: ['price-history', coinId],
    queryFn: () => fetchPriceHistory(coinId),
  });

  if (isLoading) return <p className="text-center">Loading graph...</p>;
  if (error) return <p className="text-center text-red-500">Error: {error.message}</p>;

  // Calculate min/max for dynamic range with padding
  const prices = data?.map((entry: { price: number }) => entry.price) || [];
  const minPrice = prices.length ? Math.min(...prices) : 0;
  const maxPrice = prices.length ? Math.max(...prices) : 100;
  const padding = (maxPrice - minPrice) * 0.1; // 10% padding
  const yAxisDomain = [minPrice - padding, maxPrice + padding];

  return (
    <div className="mb-4">
      <h2 className="text-xl font-semibold mb-2">Price (Past 24 Hours)</h2>
      <ResponsiveContainer width="100%" height={300}>
        <LineChart data={data}>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="timestamp" />
          <YAxis
            domain={yAxisDomain} // Dynamic range with padding
            tickFormatter={(value: number) => `$${value.toFixed(2)}`} // Format as currency
            tickCount={6} // Number of ticks
            interval="preserveStartEnd" // Ensure min/max ticks are shown
          />
          <Tooltip formatter={(value: number) => [`$${value.toFixed(2)}`, 'Price']} />
          <Line type="monotone" dataKey="price" stroke="#8884d8" />
        </LineChart>
      </ResponsiveContainer>
    </div>
  );
}