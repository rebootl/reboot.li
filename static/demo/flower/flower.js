import * as THREE from 'https://unpkg.com/three@0.160.0/build/three.module.js';

class FlowerComponent extends HTMLElement {
    constructor() {
        super();
        this.attachShadow({ mode: 'open' });
    }

    connectedCallback() {
        // Create container
        const container = document.createElement('div');
        container.style.width = '100%';
        container.style.height = '100%';
        container.style.backgroundColor = 'var(--card-background-color)';
        this.shadowRoot.appendChild(container);

        // Initialize Three.js directly since we're importing it at the top
        this.initThree(container, THREE);
    }

    initThree(container, THREE) {
        // Scene setup
        const scene = new THREE.Scene();
        scene.background = new THREE.Color(0x87CEEB);  // Sky blue background
        scene.fog = new THREE.Fog(0x87CEEB, 25, 40);  // Reduced fog density by increasing distances
        const camera = new THREE.PerspectiveCamera(75, container.clientWidth / container.clientHeight, 0.1, 1000);
        const renderer = new THREE.WebGLRenderer({ 
            antialias: true, 
            alpha: true,
            powerPreference: "high-performance"
        });
        renderer.setSize(container.clientWidth, container.clientHeight);
        container.appendChild(renderer.domElement);

        // Lighting
        const ambientLight = new THREE.AmbientLight(0xffffff, 0.8);  // Increased ambient light
        scene.add(ambientLight);

        const directionalLight = new THREE.DirectionalLight(0xffffff, 1.2);  // Increased directional light
        directionalLight.position.set(10, 15, 10);
        scene.add(directionalLight);

        // Add hemisphere light for better ambient lighting
        const hemisphereLight = new THREE.HemisphereLight(0xffffff, 0x444444, 0.8);  // Increased hemisphere light
        hemisphereLight.position.set(0, 10, 0);
        scene.add(hemisphereLight);

        // Create floor (grass)
        const floorGeometry = new THREE.PlaneGeometry(50, 50);
        const floorMaterial = new THREE.MeshPhongMaterial({ 
            color: 0x90EE90,  // Light green
            side: THREE.DoubleSide,
            flatShading: true,
            shininess: 30,
            emissive: 0x90EE90,  // Add emissive glow
            emissiveIntensity: 0.2  // Subtle glow
        });
        const floor = new THREE.Mesh(floorGeometry, floorMaterial);
        floor.rotation.x = -Math.PI / 2;  // Rotate to horizontal
        floor.position.y = -0.1;  // Slightly below ground level
        scene.add(floor);

        // Create multiple flowers at random positions
        const flowers = [];
        const flowerCount = 125;  // Number of flowers to create
        const radius = 20;  // Increased maximum distance from center
        
        for (let i = 0; i < flowerCount; i++) {
            // Random position within a circle
            const angle = Math.random() * Math.PI * 2;
            const distance = Math.random() * radius;
            const x = Math.cos(angle) * distance;
            const z = Math.sin(angle) * distance;
            const y = 0;  // All flowers start at ground level
            
            const flower = this.createFlower(THREE, x, y, z);
            flowers.push(flower);
            scene.add(flower);
        }

        // Position camera
        camera.position.set(0, 10, 25);  // Moved camera up and back
        camera.lookAt(0, 0, 0);  // Look at the center

        // Animation
        let time = 0;
        const animate = () => {
            requestAnimationFrame(animate);
            time += 0.01;
            flowers.forEach(flower => {
                flower.rotation.y = Math.sin(time + flower.position.x) * 0.5;  // Different sway for each flower
            });
            renderer.render(scene, camera);
        };
        animate();

        // Handle resize
        window.addEventListener('resize', () => {
            camera.aspect = container.clientWidth / container.clientHeight;
            camera.updateProjectionMatrix();
            renderer.setSize(container.clientWidth, container.clientHeight);
        });
    }

    createFlower(THREE, x = 0, y = 0, z = 0) {
        const flower = new THREE.Group();
        flower.position.set(x, y, z);

        // Flower color palette
        const flowerColors = [
            0xff69b4,  // Hot pink
            0xff1493,  // Deep pink
            0xffb6c1,  // Light pink
            0xffc0cb,  // Pink
            0xdda0dd,  // Plum
            0xee82ee,  // Violet
            0xda70d6,  // Orchid
            0xff00ff,  // Magenta
            0xba55d3,  // Medium orchid
            0x9370db,  // Medium purple
        ];

        // Create petals
        const petalCount = 8;
        const petalGeometry = new THREE.ConeGeometry(0.5, 1, 4);
        const petalColor = flowerColors[Math.floor(Math.random() * flowerColors.length)];
        const petalMaterial = new THREE.MeshPhongMaterial({ 
            color: petalColor,
            flatShading: true,
            emissive: petalColor,
            emissiveIntensity: 0.5
        });

        for (let i = 0; i < petalCount; i++) {
            const petal = new THREE.Mesh(petalGeometry, petalMaterial);
            const angle = (i / petalCount) * Math.PI * 2;
            petal.position.x = Math.cos(angle) * 0.5;
            petal.position.y = Math.sin(angle) * 0.5 + 3;
            petal.rotation.z = angle + Math.PI / 2;
            flower.add(petal);
        }

        // Create center disk
        const diskGeometry = new THREE.CylinderGeometry(0.3, 0.3, 0.5, 5);
        const diskMaterial = new THREE.MeshPhongMaterial({ 
            color: 0xffff00,
            flatShading: true,
            emissive: 0xffff00,
            emissiveIntensity: 0.8
        });
        const disk = new THREE.Mesh(diskGeometry, diskMaterial);
        disk.rotation.x = Math.PI / 2;
        disk.position.y = 3;
        flower.add(disk);

        // Create stem
        const stemGeometry = new THREE.CylinderGeometry(0.1, 0.1, 3, 6);
        const stemMaterial = new THREE.MeshPhongMaterial({ 
            color: 0x228b22,
            flatShading: true,
            emissive: 0x228b22,
            emissiveIntensity: 0.1
        });
        const stem = new THREE.Mesh(stemGeometry, stemMaterial);
        stem.position.y = 1.5;
        flower.add(stem);

        // Create leaves
        const leafCount = 2;
        const leafGeometry = new THREE.ConeGeometry(0.3, 0.8, 4);
        const leafMaterial = new THREE.MeshPhongMaterial({ 
            color: 0x228b22,
            flatShading: true,
            emissive: 0x228b22,
            emissiveIntensity: 0.1
        });

        for (let i = 0; i < leafCount; i++) {
            const leaf = new THREE.Mesh(leafGeometry, leafMaterial);
            const angle = (i / leafCount) * Math.PI * 2;
            leaf.position.y = 1.5 - (i * 0.5);
            leaf.position.x = Math.cos(angle) * 0.3;
            leaf.rotation.z = angle + Math.PI / 2;
            flower.add(leaf);
        }

        return flower;
    }
}

customElements.define('flower-component', FlowerComponent); 
