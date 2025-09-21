// app/layout.tsx
import type { Metadata } from 'next';
import '../styles/global.css';
import Header from '@/components/Header';
import AppProviders from '@/components/AppProviders';

export const metadata: Metadata = {
  title: 'Crypto Dashboard',
  description: 'View coin prices and AI insights',
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body>
        <AppProviders>
          <Header />
          {children}
        </AppProviders>
      </body>
    </html>
  );
}