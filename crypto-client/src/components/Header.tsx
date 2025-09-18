import Link from 'next/link'

export default function Header() {
  return (
    <header className="bg-blue-600 text-white p-4">
      <Link href="/" className="text-xl font-bold">Crypto Dashboard</Link>
    </header>
  )
}