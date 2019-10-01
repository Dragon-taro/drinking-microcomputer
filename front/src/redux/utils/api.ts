const options = {
  headers: {
    "Content-Type": "application/json"
  }
};

const url = (path: string) => `https://drinking.dragon-taro.dev/api/${path}`;

export const get = async <T>(path: string): Promise<T> => {
  // エラーハンドリングはめんどいしなし
  const resp = await fetch(url(path), options);
  const json = (await resp.json()) as T;

  return json;
};

export const post = async <T>(path: string): Promise<T> => {
  // エラーハンドリングはめんどいしなし
  // content-type要らんかったw
  const postOptions = { ...options, method: "POST" };
  const resp = await fetch(url(path), postOptions);
  const json = (await resp.json()) as T;

  return json;
};

export const patch = async <T>(path: string): Promise<T> => {
  // エラーハンドリングはめんどいしなし
  // content-type要らんかったw
  const postOptions = { ...options, method: "PATCH" };
  const resp = await fetch(url(path), postOptions);
  const json = (await resp.json()) as T;

  return json;
};
