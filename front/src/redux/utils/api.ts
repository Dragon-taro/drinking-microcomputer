const options = {
  headers: {
    "Content-Type": "application/json"
  }
};

const url = (path: string) => `https://drinking.dragon-taro.dev/${path}`;

export const get = async <T>(path: string): Promise<T> => {
  // エラーハンドリングはめんどいしなし
  const resp = await fetch(url(path), options);
  const json = (await resp.json()) as T;

  return json;
};
