import Head from 'next/head';
import { useState, FormEvent } from 'react';
import fetch from '../lib/fetch';

export default function Home() {
  const [auth, setAuth] = useState(false);

  const [form, setForm] = useState({
    username: '',
    password: '',
  });

  const onSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    const result = await fetch<{ accessToken: string }>('/auth/token', {
      body: JSON.stringify(form),
      method: 'POST',
    });

    localStorage.setItem('token', result.accessToken);

    setAuth(true);
  };

  const meRequest = async () => {
    const result = await fetch('/auth/me', {
      method: 'GET',
    });

    alert(JSON.stringify(result));
  };

  const onChange = (event: React.FormEvent<EventTarget>) => {
    const name = (event.target as HTMLInputElement).name;
    const value = (event.target as HTMLInputElement).value;
    setForm((old) => ({ ...old, [name]: value }));
  };

  const silentRefresh = async () => {
    const result = await fetch<{ accessToken: string }>('/auth/refresh', {
      method: 'POST',
    });

    localStorage.setItem('token', result.accessToken);
  };

  return (
    <div className="container">
      <Head>
        <title>Monorepo app</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main>
        <form className="box" onSubmit={onSubmit}>
          <input name="username" onChange={onChange} />
          <input name="password" onChange={onChange} />
          <button type="submit">SUBMIT</button>
        </form>

        <button type="button" onClick={silentRefresh}>
          Silent refresh
        </button>
        <button onClick={meRequest}>Me</button>
      </main>

      <footer>Footer</footer>

      <style jsx>{`
        .box {
          display: flex;
          justify-content: center;
          align-items: center;
          flex-direction: column;
        }

        .box input {
          margin-top: 32px;
        }
      `}</style>

      <style jsx global>{`
        main {
          height: 100vh;
        }

        html,
        body {
          height: 100vh;
          padding: 0;
          margin: 0;
          font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Roboto,
            Oxygen, Ubuntu, Cantarell, Fira Sans, Droid Sans, Helvetica Neue,
            sans-serif;
        }

        * {
          box-sizing: border-box;
        }
      `}</style>
    </div>
  );
}
