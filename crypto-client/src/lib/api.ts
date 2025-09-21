

export async function fetchPriceHistory(coinId: string) {
  
  const endpoint = process.env.NEXT_PUBLIC_COIN_ENDPOINT

  const res = await fetch(`${endpoint}/coins/${coinId}/market_chart?vs_currency=usd&days=1`, {
    next: { revalidate: 300 },
  });
  if (!res.ok) throw new Error('Failed to fetch price history');
  const data = await res.json();
  return data.prices.map(([timestamp, price]: [number, number]) => ({
    timestamp: new Date(timestamp).toLocaleTimeString(),
    price,
  }));
}

export async function fetchAIInsights(coinId: string) {
  const res = await fetch(`https://data-ingestion-abcdef-ue.a.run.app/ingest/${coinId}`)  // Your backend API
  if (!res.ok) throw new Error('Failed to fetch AI insights')
  return res.json()
}