import { useRef, useEffect } from 'react';

const Canvas = ({draw, height, width}) => {
    const canvas = useRef();

    useEffect(() => {
        const context = canvas.current.getContext('2d');
        draw(context);
    });

    return (
        <canvas ref={canvas} height={height} width={width} />
    );
}

export default Canvas;
