import React from 'react';
import useSWR from 'swr';

// TODO: Move Unit typedef to a separate ApiResponse definitions file

/**
 * @typedef {object} Unit
 *
 * @property {string} ID
 * @property {string} MDNS
 */

/**
 * @return {React.ReactElement}
 */
export default function ListUnits() {
	const { data: resp, error } = useSWR('/api/get-paired-units');

	if (error || resp?.error) return <div>failed to load</div>;
	if (!resp) return <div>loading...</div>;
	if (!resp.data.length) {
		return <div>No data</div>;
	}

	return resp.data.map((/** @type {Unit} */ unit) => (
		<div key={`${unit.ID}`}>{`${unit.MDNS}`}</div>
	));
}

