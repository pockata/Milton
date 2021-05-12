// Next.js API route support: https://nextjs.org/docs/api-routes/introduction

/** @typedef {import('next').NextApiResponse} NextApiResponse */
/** @typedef {import('next').NextApiRequest} NextApiRequest */

/**
 * @param {NextApiRequest} _req
 * @param {NextApiResponse} res
 */

export default (_req, res) => {
	res.status(200).json({ name: 'John Doe' });
};

