'use client'  // Client-side for interactivity
import CoinList from '@/components/CoinList'
import { useQuery, QueryClientProvider, QueryClient} from '@tanstack/react-query'
import * as dotenv from 'dotenv';


const queryClient = new QueryClient();
dotenv.config();

export default function Home() {
  return (
    <QueryClientProvider client={queryClient}>
      <main className="container mx-auto p-4">
        <h1 className="text-3xl font-bold mb-4">Crypto Dashboard</h1>
        <CoinList />
      </main>
    </QueryClientProvider>
  )
}