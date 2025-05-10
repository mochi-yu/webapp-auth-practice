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

  await apiClient
    .get(`/signin/github?code=${code}`, { cache: "no-store" })
    .catch((err) => {
      alert("ログインに失敗しました。少し待ってからやり直してください。");
      console.error(err);
    });
  redirect("/");
}
