import { useQuery } from '@tanstack/react-query'
import { fetchAIInsights } from '@/lib/api'

export default function AIInsights({ coinId }: { coinId: string }) {
  const { data, isLoading } = useQuery({
    queryKey: ['ai-insights', coinId],
    queryFn: () => fetchAIInsights(coinId),
  })

  if (isLoading) return <p>Loading AI insights...</p>

  return (
    <div className="mt-4 p-4 border rounded">
      <h2 className="text-2xl font-bold">AI Insights</h2>
      <p>{data?.investment_insight}</p>
      {/* Display articles or other data if needed */}
    </div>
  )
}