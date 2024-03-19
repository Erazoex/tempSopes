import React, { useRef, useEffect } from "react";
import { DataSet, Network } from "vis-network/standalone";
import "vis-network/styles/vis-network.css";

const Tree = ({ process }) => {
  const containerRef = useRef(null);

  useEffect(() => {
    const nodes = new DataSet([
      { id: 1, label: "Raíz" },
      { id: 2, label: "Nodo 1" },
      { id: 3, label: "Nodo 2" },
      { id: 4, label: "Nodo 3" }
    ]);

    const edges = new DataSet([
      { from: 1, to: 2 },
      { from: 1, to: 3 },
      { from: 1, to: 4 }
    ]);

    const data = {
      nodes: nodes,
      edges: edges
    };

    const options = {
      layout: {
        hierarchical: {
          direction: "UD",
          sortMethod: "directed"
        }
      }
    };

    const network = new Network(containerRef.current, data, options);

    return () => {
      network.destroy();
    };
  }, []);

  return <div ref={containerRef} style={{ width: "600px", height: "400px" }} />;
};

export default Tree;