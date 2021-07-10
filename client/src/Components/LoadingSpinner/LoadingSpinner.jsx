import './LoadingSpinner.scss';

/**
 *
 * @param {Object} props
 * @param {string} radius radius of spinner in css <length> format
 * @param {string} strokeWidth stroke of spinner in css <length> format
 * @returns
 */
const LoadingSpinner = ({
	radius,
	strokeWidth,
	strokeColor = 'red',
	fillColor = 'transparent',
}) => {
	return (
		<svg
			style={{
				'--radius': radius,
				'--stroke-width': strokeWidth,
			}}
			className='loading-spinner'
		>
			<circle stroke={strokeColor} fill={fillColor} strokeWidth={strokeWidth} />
		</svg>
	);
};

export default LoadingSpinner;
