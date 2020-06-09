import fetch from 'isomorphic-unfetch';

export default async function <JSON = any>(
  input: RequestInfo,
  init?: RequestInit
) {
  const headers = init.headers ? init.headers : {};
  headers['Content-Type'] = 'application/json';

  if (typeof window !== 'undefined') {
    const token = localStorage.getItem('token');
    if (token) {
      headers['Authorization'] = `Bearer ${token}`;
    }
  }

  const res = await fetch('http://localhost:8080' + input, {
    ...init,
    credentials: 'include',
    mode: 'cors',
    headers: new Headers({ ...headers }),
  });
  return res.json() as Promise<JSON>;
}
