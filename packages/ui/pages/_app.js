import "normalize.css"
import '../styles/globals.css'

import { SWRConfig } from 'swr'

function MyApp({ Component, pageProps }) {
  let fetcher = (resource, init) => fetch(resource, init).then(res => res.json())
  return (
    <SWRConfig value={{fetcher}} >
      <Component {...pageProps} />
    </SWRConfig>
  )
}

export default MyApp
