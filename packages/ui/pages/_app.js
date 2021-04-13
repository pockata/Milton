import "normalize.css"
import '../styles/globals.scss'

import { SWRConfig } from 'swr'
import config from '../config'

function MyApp({ Component, pageProps }) {
  let fetcher = (resource, init) => fetch(config.apiURL + resource, init).then(res => res.json())
  return (
    <SWRConfig value={{fetcher}} >
      <Component {...pageProps} />
    </SWRConfig>
  )
}

export default MyApp
