import CountryList from '../components/CountryList';

export default function Page() {
  return (
    <>
      <h2>国一覧</h2>
      <p>GraphQL API を使用して国のデータを取得しています。</p>
      <CountryList />
    </>
  );
}