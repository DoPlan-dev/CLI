import { useEffect, useRef } from 'react';

interface Point {
    x: number;
    y: number;
    z: number;
    baseX: number;
    baseY: number;
    baseZ: number;
    phase: number;
}



export default function AnimatedBackground() {
    const canvasRef = useRef<HTMLCanvasElement>(null);
    const mouseRef = useRef({ x: 0, y: 0 });
    const containerRef = useRef<HTMLDivElement>(null);

    useEffect(() => {
        const canvas = canvasRef.current;
        if (!canvas) return;

        const ctx = canvas.getContext('2d');
        if (!ctx) return;

        let animationFrameId: number;
        let time = 0;

        // Configuration
        const GRID_ROWS = 30;
        const GRID_COLS = 30;
        const GRID_SPACING = 80;
        const WAVE_HEIGHT = 80;
        const MOUSE_INFLUENCE_RADIUS = 350;

        // State
        let gridPoints: Point[] = [];

        const resizeCanvas = () => {
            if (containerRef.current && canvas) {
                canvas.width = containerRef.current.offsetWidth;
                canvas.height = containerRef.current.offsetHeight;
                initScene();
            }
        };

        const initScene = () => {
            gridPoints = [];

            // Initialize Grid (The "Structure")
            const startX = -(GRID_COLS * GRID_SPACING) / 2;
            const startZ = -(GRID_ROWS * GRID_SPACING) / 2;

            for (let i = 0; i < GRID_COLS; i++) {
                for (let j = 0; j < GRID_ROWS; j++) {
                    const x = startX + i * GRID_SPACING;
                    const z = startZ + j * GRID_SPACING;
                    gridPoints.push({
                        x, y: 0, z,
                        baseX: x, baseY: 0, baseZ: z,
                        phase: Math.random() * Math.PI * 2
                    });
                }
            }
        };

        const project = (x: number, y: number, z: number) => {
            const fov = 800;
            const scale = fov / (fov + z + 1000); // Offset camera back
            const x2d = x * scale + canvas.width / 2;
            const y2d = y * scale + canvas.height / 2;
            return { x: x2d, y: y2d, scale };
        };

        const draw = () => {
            ctx.clearRect(0, 0, canvas.width, canvas.height);
            time += 0.015;

            // Get current theme color
            const style = getComputedStyle(document.documentElement);
            const foreground = style.getPropertyValue('--foreground').trim();

            // Camera rotation
            const angleX = 0.3; // Tilt

            // 1. Update and Draw Grid
            gridPoints.forEach(p => {
                // Wave Motion
                const dist = Math.sqrt(p.baseX * p.baseX + p.baseZ * p.baseZ);
                p.y = Math.sin(dist * 0.01 - time * 2) * WAVE_HEIGHT + Math.sin(p.baseX * 0.02 + time) * 20;

                // Rotate
                const rx = p.x;
                const ry = p.y * Math.cos(angleX) - p.z * Math.sin(angleX);
                const rz = p.y * Math.sin(angleX) + p.z * Math.cos(angleX);

                // Project
                const proj = project(rx, ry, rz);

                // Mouse Interaction (2D)
                const dx = proj.x - (mouseRef.current.x * canvas.width / 2 + canvas.width / 2);
                const dy = proj.y - (mouseRef.current.y * canvas.height / 2 + canvas.height / 2);
                const dist2d = Math.sqrt(dx * dx + dy * dy);

                let alpha = Math.max(0, 1 - (rz + 1000) / 2000) * 0.5; // Base opacity
                let scale = proj.scale;

                if (dist2d < MOUSE_INFLUENCE_RADIUS) {
                    alpha += (1 - dist2d / MOUSE_INFLUENCE_RADIUS) * 0.5;
                    scale *= 1 + (1 - dist2d / MOUSE_INFLUENCE_RADIUS);
                }

                if (alpha > 0.01) {
                    ctx.beginPath();
                    ctx.arc(proj.x, proj.y, 2.5 * scale, 0, Math.PI * 2);
                    ctx.fillStyle = `hsl(${foreground} / ${alpha})`;
                    ctx.fill();
                }
            });

            animationFrameId = requestAnimationFrame(draw);
        };

        const handleMouseMove = (e: MouseEvent) => {
            const rect = canvas.getBoundingClientRect();
            mouseRef.current = {
                x: (e.clientX - rect.left - canvas.width / 2) / (canvas.width / 2),
                y: (e.clientY - rect.top - canvas.height / 2) / (canvas.height / 2)
            };
        };

        window.addEventListener('resize', resizeCanvas);
        window.addEventListener('mousemove', handleMouseMove);

        resizeCanvas();
        draw();

        return () => {
            window.removeEventListener('resize', resizeCanvas);
            window.removeEventListener('mousemove', handleMouseMove);
            cancelAnimationFrame(animationFrameId);
        };
    }, []);

    return (
        <div ref={containerRef} className="absolute inset-0 overflow-hidden pointer-events-none">
            <canvas
                ref={canvasRef}
                className="absolute inset-0"
                style={{ opacity: 0.8 }}
            />
            <div className="absolute inset-0 bg-gradient-to-b from-transparent via-transparent to-background/50" />
        </div>
    );
}
