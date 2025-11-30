// Three.js 3D Glitch Background - "The Future of AI Development Glitches"
(function() {
    const container = document.getElementById('threejs-container');
    if (!container) return;

    // Scene setup
    const scene = new THREE.Scene();
    const camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000);
    const renderer = new THREE.WebGLRenderer({ 
        antialias: true,
        alpha: true,
        powerPreference: "high-performance"
    });
    
    renderer.setSize(window.innerWidth, window.innerHeight);
    renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2));
    container.appendChild(renderer.domElement);

    // Glitch effect uniforms
    const glitchUniforms = {
        time: { value: 0 },
        intensity: { value: 0.5 },
        color1: { value: new THREE.Color(0x89b4fa) }, // Yuki7
        color2: { value: new THREE.Color(0xf38ba8) }, // Yuki11
        color3: { value: new THREE.Color(0x89dceb) }  // Yuki9
    };

    // Create glitch shader material
    const glitchMaterial = new THREE.ShaderMaterial({
        uniforms: glitchUniforms,
        vertexShader: `
            uniform float time;
            uniform float intensity;
            varying vec2 vUv;
            varying vec3 vPosition;
            
            void main() {
                vUv = uv;
                vPosition = position;
                
                vec3 pos = position;
                
                // Glitch distortion
                float glitch = sin(time * 10.0) * intensity;
                pos.x += sin(pos.y * 10.0 + time * 5.0) * glitch * 0.1;
                pos.y += cos(pos.x * 10.0 + time * 5.0) * glitch * 0.1;
                
                gl_Position = projectionMatrix * modelViewMatrix * vec4(pos, 1.0);
            }
        `,
        fragmentShader: `
            uniform float time;
            uniform float intensity;
            uniform vec3 color1;
            uniform vec3 color2;
            uniform vec3 color3;
            varying vec2 vUv;
            varying vec3 vPosition;
            
            void main() {
                vec2 uv = vUv;
                
                // Glitch effect
                float glitch = sin(time * 20.0) * intensity;
                float scanline = sin(uv.y * 800.0 + time * 10.0) * 0.5 + 0.5;
                
                // Color channel separation (RGB shift)
                float r = sin(uv.x * 10.0 + time * 5.0) * glitch * 0.1;
                float g = sin(uv.x * 10.0 + time * 5.0 + 2.0) * glitch * 0.1;
                float b = sin(uv.x * 10.0 + time * 5.0 + 4.0) * glitch * 0.1;
                
                // Digital noise
                float noise = fract(sin(dot(uv, vec2(12.9898, 78.233)) * time) * 43758.5453);
                
                // Combine colors with glitch
                vec3 color = mix(color1, color2, sin(time * 2.0) * 0.5 + 0.5);
                color = mix(color, color3, sin(time * 3.0) * 0.5 + 0.5);
                
                // Apply effects
                color.r += r * intensity;
                color.g += g * intensity;
                color.b += b * intensity;
                color += noise * 0.1 * intensity;
                
                // Scanline effect
                color *= (0.9 + scanline * 0.1);
                
                // Fade edges
                float edge = smoothstep(0.0, 0.1, uv.x) * smoothstep(1.0, 0.9, uv.x) *
                            smoothstep(0.0, 0.1, uv.y) * smoothstep(1.0, 0.9, uv.y);
                
                gl_FragColor = vec4(color * edge, 0.15);
            }
        `,
        transparent: true,
        side: THREE.DoubleSide
    });

    // Create geometry for glitch effect
    const geometry = new THREE.PlaneGeometry(20, 20, 32, 32);
    const glitchPlane = new THREE.Mesh(geometry, glitchMaterial);
    scene.add(glitchPlane);

    // Add wireframe grid
    const gridHelper = new THREE.GridHelper(20, 20, 0x89b4fa, 0x313244);
    gridHelper.material.opacity = 0.2;
    gridHelper.material.transparent = true;
    scene.add(gridHelper);

    // Add floating geometric shapes (glitched)
    const shapes = [];
    const shapeCount = 8;
    
    for (let i = 0; i < shapeCount; i++) {
        const geometry = new THREE.BoxGeometry(0.5, 0.5, 0.5);
        const material = new THREE.MeshBasicMaterial({
            color: i % 3 === 0 ? 0x89b4fa : i % 3 === 1 ? 0xf38ba8 : 0x89dceb,
            wireframe: true,
            transparent: true,
            opacity: 0.3
        });
        const shape = new THREE.Mesh(geometry, material);
        
        shape.position.set(
            (Math.random() - 0.5) * 10,
            (Math.random() - 0.5) * 10,
            (Math.random() - 0.5) * 5
        );
        
        shape.userData = {
            speed: Math.random() * 0.02 + 0.01,
            rotationSpeed: Math.random() * 0.02 + 0.01
        };
        
        shapes.push(shape);
        scene.add(shape);
    }

    // Add particles
    const particleGeometry = new THREE.BufferGeometry();
    const particleCount = 200;
    const positions = new Float32Array(particleCount * 3);
    
    for (let i = 0; i < particleCount * 3; i++) {
        positions[i] = (Math.random() - 0.5) * 20;
    }
    
    particleGeometry.setAttribute('position', new THREE.BufferAttribute(positions, 3));
    const particleMaterial = new THREE.PointsMaterial({
        color: 0x89b4fa,
        size: 0.05,
        transparent: true,
        opacity: 0.6
    });
    const particles = new THREE.Points(particleGeometry, particleMaterial);
    scene.add(particles);

    // Mouse interaction
    let mouseX = 0;
    let mouseY = 0;
    let targetX = 0;
    let targetY = 0;
    
    document.addEventListener('mousemove', (e) => {
        targetX = (e.clientX / window.innerWidth) * 2 - 1;
        targetY = -(e.clientY / window.innerHeight) * 2 + 1;
    });

    // Camera position
    camera.position.z = 5;

    // Animation loop
    let time = 0;
    function animate() {
        requestAnimationFrame(animate);
        
        time += 0.01;
        
        // Smooth mouse tracking
        mouseX += (targetX - mouseX) * 0.05;
        mouseY += (targetY - mouseY) * 0.05;
        
        // Update glitch uniforms
        glitchUniforms.time.value = time;
        glitchUniforms.intensity.value = 0.3 + Math.sin(time * 2) * 0.2;
        
        // Rotate camera based on mouse
        camera.position.x += (mouseX * 2 - camera.position.x) * 0.05;
        camera.position.y += (mouseY * 2 - camera.position.y) * 0.05;
        camera.lookAt(0, 0, 0);
        
        // Animate shapes
        shapes.forEach((shape, i) => {
            shape.rotation.x += shape.userData.rotationSpeed;
            shape.rotation.y += shape.userData.rotationSpeed;
            shape.position.y += Math.sin(time * 2 + i) * 0.01;
            shape.position.x += Math.cos(time * 2 + i) * 0.01;
            
            // Glitch effect on shapes
            if (Math.random() > 0.98) {
                shape.position.x += (Math.random() - 0.5) * 0.5;
                shape.position.y += (Math.random() - 0.5) * 0.5;
            }
        });
        
        // Rotate particles
        particles.rotation.y += 0.001;
        
        // Random glitch intensity spikes
        if (Math.random() > 0.95) {
            glitchUniforms.intensity.value = 1.0;
        }
        
        renderer.render(scene, camera);
    }

    // Handle resize
    function handleResize() {
        camera.aspect = window.innerWidth / window.innerHeight;
        camera.updateProjectionMatrix();
        renderer.setSize(window.innerWidth, window.innerHeight);
    }
    
    window.addEventListener('resize', handleResize);
    
    // Update colors based on theme
    function updateColors() {
        const isLight = document.documentElement.getAttribute('data-theme') === 'light';
        if (isLight) {
            glitchUniforms.color1.value.setHex(0x89b4fa);
            glitchUniforms.color2.value.setHex(0xf38ba8);
            glitchUniforms.color3.value.setHex(0x89dceb);
        } else {
            glitchUniforms.color1.value.setHex(0x89b4fa);
            glitchUniforms.color2.value.setHex(0xf38ba8);
            glitchUniforms.color3.value.setHex(0x89dceb);
        }
    }
    
    // Watch for theme changes
    const observer = new MutationObserver(updateColors);
    observer.observe(document.documentElement, { attributes: true, attributeFilter: ['data-theme'] });
    
    animate();
})();

