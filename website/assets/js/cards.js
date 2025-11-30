// Three.js Card Animations
function initThreeJSCard(canvasId) {
    const canvas = document.getElementById(canvasId);
    if (!canvas) return;

    const scene = new THREE.Scene();
    const camera = new THREE.PerspectiveCamera(50, canvas.clientWidth / canvas.clientHeight, 0.1, 1000);
    const renderer = new THREE.WebGLRenderer({ 
        canvas: canvas, 
        alpha: true,
        antialias: true 
    });
    
    renderer.setSize(canvas.clientWidth, canvas.clientHeight);
    renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2));

    // Create multiple geometric shapes for a more interesting scene
    const group = new THREE.Group();
    
    // Main wireframe cube
    const geometry1 = new THREE.BoxGeometry(1.5, 1.5, 1.5);
    const material1 = new THREE.MeshBasicMaterial({ 
        color: 0xffffff,
        wireframe: true,
        opacity: 0.4,
        transparent: true
    });
    const cube1 = new THREE.Mesh(geometry1, material1);
    group.add(cube1);
    
    // Smaller inner cube
    const geometry2 = new THREE.BoxGeometry(0.8, 0.8, 0.8);
    const material2 = new THREE.MeshBasicMaterial({ 
        color: 0xcccccc,
        wireframe: true,
        opacity: 0.3,
        transparent: true
    });
    const cube2 = new THREE.Mesh(geometry2, material2);
    group.add(cube2);
    
    // Add some particles
    const particlesGeometry = new THREE.BufferGeometry();
    const particlesCount = 20;
    const posArray = new Float32Array(particlesCount * 3);
    
    for (let i = 0; i < particlesCount * 3; i++) {
        posArray[i] = (Math.random() - 0.5) * 4;
    }
    
    particlesGeometry.setAttribute('position', new THREE.BufferAttribute(posArray, 3));
    const particlesMaterial = new THREE.PointsMaterial({
        size: 0.05,
        color: 0xffffff,
        transparent: true,
        opacity: 0.6
    });
    const particles = new THREE.Points(particlesGeometry, particlesMaterial);
    group.add(particles);
    
    scene.add(group);
    camera.position.z = 4;

    // Mouse tracking for interactive rotation
    let mouseX = 0;
    let mouseY = 0;
    let targetRotationX = 0;
    let targetRotationY = 0;
    
    const card = canvas.closest('.threejs-card');
    if (card) {
        card.addEventListener('mousemove', (e) => {
            const rect = card.getBoundingClientRect();
            mouseX = ((e.clientX - rect.left) / rect.width) * 2 - 1;
            mouseY = -((e.clientY - rect.top) / rect.height) * 2 + 1;
            targetRotationY = mouseX * 0.5;
            targetRotationX = mouseY * 0.5;
        });
        
        card.addEventListener('mouseenter', () => {
            material1.opacity = 0.6;
            material2.opacity = 0.5;
            particlesMaterial.opacity = 0.8;
        });
        
        card.addEventListener('mouseleave', () => {
            material1.opacity = 0.4;
            material2.opacity = 0.3;
            particlesMaterial.opacity = 0.6;
            targetRotationX = 0;
            targetRotationY = 0;
        });
    }

    // Animation
    function animate() {
        requestAnimationFrame(animate);
        
        // Smooth rotation towards target
        group.rotation.y += (targetRotationY - group.rotation.y) * 0.1;
        group.rotation.x += (targetRotationX - group.rotation.x) * 0.1;
        
        // Continuous slow rotation
        group.rotation.z += 0.005;
        
        // Rotate cubes independently
        cube1.rotation.x += 0.01;
        cube1.rotation.y += 0.01;
        cube2.rotation.x -= 0.015;
        cube2.rotation.y -= 0.015;
        
        // Rotate particles
        particles.rotation.y += 0.002;
        
        renderer.render(scene, camera);
    }

    // Handle resize
    function handleResize() {
        camera.aspect = canvas.clientWidth / canvas.clientHeight;
        camera.updateProjectionMatrix();
        renderer.setSize(canvas.clientWidth, canvas.clientHeight);
    }
    
    window.addEventListener('resize', handleResize);
    
    animate();
}

// Initialize Three.js cards when page loads
document.addEventListener('DOMContentLoaded', () => {
    initThreeJSCard('threejs-canvas-1');
});

// Card selection handler
document.querySelectorAll('.card-option').forEach(option => {
    option.addEventListener('click', () => {
        const design = option.getAttribute('data-design');
        console.log('Selected design:', design);
        // You can add logic here to apply the selected design to the main page
    });
});

