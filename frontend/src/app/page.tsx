export default function Home() {
  return (
    <div>
      <h1>Home</h1>
      <a
        href={`https://github.com/login/oauth/authorize?client_id=${process.env.NEXT_PUBLIC_GITHUB_CLIENT_ID}`}
        className="bg-black cursor-pointer text-white px-4 py-2 rounded-md"
      >
        GitHub でログイン
      </a>
    </div>
  );
}
