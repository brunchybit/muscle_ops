export function App() {
    const canvasRef = React.useRef(null);

    return (
        <div>
            <h1>Muscle Ops</h1>
            <canvas width={1024}
                    height={800}
                    onClick={e => {
                        const canvas = canvasRef.current
                        const ctx = canvas.getContext('2d');
                        alert(e.clientX)
                    }}/>
        </div>
    );
}
