// fetchをラップしたAPIクライアント
export const apiClient = {
  async request(
    path: string,
    method: string,
    body?: unknown,
    options?: RequestInit,
  ) {
    options = {
      method,
      headers: body
        ? {
            "Content-Type": "application/json",
          }
        : undefined,
      body: body ? JSON.stringify(body) : undefined,
      ...options,
    };

    const res = await fetch(
      `${process.env.NEXT_PUBLIC_API_URL}${path}`,
      options,
    );
    return res;
  },

  async get(path: string, options?: RequestInit) {
    return this.request(path, "GET", undefined, options);
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
