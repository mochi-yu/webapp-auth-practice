export default function Home() {
  return (
    <div>
      <h1>Home</h1>
      <a
        href={`https://github.com/login/oauth/authorize?client_id=${process.env.NEXT_PUBLIC_GITHUB_CLIENT_ID}`}
      >
        Sign in with Github
      </a>
    </div>
  );
}
