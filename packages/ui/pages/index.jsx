import Head from 'next/head';
import React from 'react';
// import styles from '../styles/Home.module.css'

import ListUnits from '../components/ListUnits';

import Layout from '../layouts/main';

export default function Home() {
	return (
		<Layout>
			<Head>
				<title>Въй Лек</title>
				<link rel="icon" href="/favicon.ico" />
			</Head>
			<ListUnits/>
		</Layout>
	);
}

