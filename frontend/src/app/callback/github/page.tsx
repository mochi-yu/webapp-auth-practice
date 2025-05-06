import { apiClient } from "@/libs/api/client";
import { redirect } from "next/navigation";

type Props = {
  searchParams: {
    code: string;
  };
};

export default async function GithubCallbackPage({ searchParams }: Props) {
  // クエリパラメータからcodeを取得
  const { code } = await searchParams;

  // codeがない場合はリダイレクト
  if (!code) {
    redirect("/");
  }

  const res = await apiClient.get(`/signin/github?code=${code}`);
  console.log(res);
  redirect("/");

  return (
    <div>
      <h1>GithubCallbackPage</h1>
    </div>
  );
}
