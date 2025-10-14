'use client';
import './globals.css';
import { Provider } from 'urql';
import { client } from '../lib/urqlClient';
import Header from '../components/Header';

export const metadata = {
  title: 'Next.js + GraphQL サンプル',
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="ja">
      <body>
        <Provider value={client}>
          <Header />
          <main style={{ padding: 16 }}>{children}</main>
        </Provider>
      </body>
    </html>
  );
}