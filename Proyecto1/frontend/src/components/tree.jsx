import React, { useEffect, useRef } from 'react';
import { Network, DataSet } from 'vis-network/standalone/esm/vis-network.min.js';

const ProcessTree = ({ process }) => {
  const containerRef = useRef(null);
  const networkRef = useRef(null);

  useEffect(() => {
    if (!process || !containerRef.current) return;

    // Crear un DataSet para los nodos y aristas
    const nodes = new DataSet();
    const edges = new DataSet();

    // Agregar el nodo raíz
    nodes.add({ id: process.pid, label: process.name });

    // Función recursiva para agregar nodos y aristas
    const addNodesAndEdges = (parent, children) => {
      children.forEach(child => {
        nodes.add({ id: child.pid, label: child.name });
        edges.add({ from: parent.pid, to: child.pid });
        if (child.child && child.child.length > 0) {
          addNodesAndEdges(child, child.child);
        }
      });
    };

    // Agregar nodos y aristas
    if (process.child && process.child.length > 0) {
      addNodesAndEdges(process, process.child);
    }

    // Configurar la red
    const data = { nodes, edges };
    const options = {
      layout: {
        hierarchical: {
          direction: "UD", // De arriba hacia abajo
          sortMethod: "directed", // Ordena los nodos según la dirección de las aristas
          levelSeparation: 150, // Separación entre niveles del árbol
          nodeSpacing: 100 // Espaciado entre nodos
        }
      }
    };
    networkRef.current = new Network(containerRef.current, data, options);

    return () => {
      if (networkRef.current) {
        networkRef.current.destroy();
        networkRef.current = null;
      }
    };
  }, [process]);

  return <div ref={containerRef} style={{ width: '100%', height: '70vh' }} />;
};

export default ProcessTree;
