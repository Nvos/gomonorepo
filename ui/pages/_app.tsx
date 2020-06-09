import { AppProps, AppInitialProps, AppContext } from 'next/app';
import fetch from '../lib/fetch';
import { useEffect } from 'react';

type Props = AppProps & { accessToken: string };

function CustomApp({ Component, pageProps, accessToken, router }: Props) {
  // useEffect(() => {
  //   if (typeof window !== `undefined`) {
  //     if (!accessToken) {
  //       router.push('/login');
  //     }

  //     localStorage.setItem('token', accessToken);
  //   }
  // }, []);

  return <Component {...pageProps} />;
}

const silentRefresh = async ({ cookie }: { cookie: string }) => {
  try {
    const result = await fetch<{ accessToken: string }>('/auth/refresh', {
      method: 'POST',
      headers: {
        cookie: cookie,
      },
    });

    return result?.accessToken ?? undefined;
  } catch (e) {
    return undefined;
  }
};

CustomApp.getInitialProps = async ({ ctx }: AppContext) => {
  console.log(ctx.req.headers.cookie);

  const token = await silentRefresh({ cookie: ctx.req.headers.cookie });

  console.log('initial props', token);
  return {
    accessToken: token,
  };
};

export default CustomApp;
