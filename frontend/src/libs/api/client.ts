// fetchをラップしたAPIクライアント
export const apiClient = {
  async request(path: string, method: string, body?: unknown) {
    const options: RequestInit = {
      method,
      headers: body
        ? {
            "Content-Type": "application/json",
          }
        : undefined,
      body: body ? JSON.stringify(body) : undefined,
    };

    const res = await fetch(
      `${process.env.NEXT_PUBLIC_API_URL}${path}`,
      options,
    );
    return res;
  },

  async get(path: string) {
    return this.request(path, "GET");
  },

  async post(path: string, body?: unknown) {
    return this.request(path, "POST", body);
  },

  async put(path: string, body?: unknown) {
    return this.request(path, "PUT", body);
  },

  async delete(path: string) {
    return this.request(path, "DELETE");
  },

  async patch(path: string, body?: unknown) {
    return this.request(path, "PATCH", body);
  },
};
