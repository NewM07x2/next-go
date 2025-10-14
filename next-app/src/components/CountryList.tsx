'use client';
import { useQuery } from 'urql';

const COUNTRIES_QUERY = `
  query {
    countries {
      code
      name
      emoji
      continent {
        name
      }
    }
  }
`;

export default function CountryList() {
  const [{ data, fetching, error }] = useQuery({ query: COUNTRIES_QUERY });

  if (fetching) return <p>読み込み中...</p>;
  if (error) return <p>エラー: {error.message}</p>;

  return (
    <ul style={{ listStyle: 'none', padding: 0 }}>
      {data.countries.map((country: any) => (
        <li key={country.code} style={{ padding: 8, borderBottom: '1px solid #f0f0f0' }}>
          <strong>{country.name}</strong> {country.emoji} <span style={{ color: '#666' }}>({country.continent.name})</span>
        </li>
      ))}
    </ul>
  );
}