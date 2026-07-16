#!/usr/bin/env node

/**
 * This script reseta o projeto para uma estrutura moderna dentro de /src.
 * Ele move ou deleta as pastas antigas e cria src/app, src/components, src/hooks, src/constants e src/scripts.
 * Após rodar, remova este script e o comando do package.json.
 */

const fs = require("fs");
const path = require("path");
const readline = require("readline");

const root = process.cwd();
const oldDirs = ["app", "components", "hooks", "constants", "scripts", "features", "services", "styles", "utils"];
const exampleDir = "app-example";
const srcDir = path.join(root, "src");
const newDirs = [
  path.join("src", "app"),
  path.join("src", "components"),
  path.join("src", "hooks"),
  path.join("src", "constants"),
  path.join("src", "features"),
  path.join("src", "services"),
  path.join("src", "styles"),
  path.join("src", "utils"),
];
const exampleDirPath = path.join(root, exampleDir);

const indexContent = `import SolarFacilApp from '../components/SolarFacilApp';

export default function Index() {
  return <SolarFacilApp />;
}
`;

const layoutContent = `import { Stack } from "expo-router";

export default function RootLayout() {
  return <Stack />;
}
`;

const componentContent = `import React from 'react';
import { View, Text, TextInput, TouchableOpacity, StyleSheet } from 'react-native';

const SolarFacilApp: React.FC = () => (
  <View style={styles.container}>
    <Text style={styles.title}>Solar Fácil</Text>
    <Text style={styles.subtitle}>Entrar ou Criar Conta</Text>
    <TextInput style={styles.input} placeholder="Email" keyboardType="email-address" />
    <TextInput style={styles.input} placeholder="Senha" secureTextEntry />
    <TouchableOpacity style={styles.button}>
      <Text style={styles.buttonText}>Entrar</Text>
    </TouchableOpacity>
    <TouchableOpacity style={[styles.button, styles.secondaryButton]}>
      <Text style={styles.buttonText}>Criar Conta</Text>
    </TouchableOpacity>
  </View>
);

export default SolarFacilApp;

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#F2FBE0',
    alignItems: 'center',
    justifyContent: 'center',
    padding: 20
  },
  title: {
    fontSize: 32,
    fontWeight: 'bold',
    color: '#2D572C',
    marginBottom: 10
  },
  subtitle: {
    fontSize: 18,
    color: '#555',
    marginBottom: 20
  },
  input: {
    width: '100%',
    height: 50,
    backgroundColor: '#fff',
    borderRadius: 10,
    paddingHorizontal: 15,
    marginBottom: 15,
    borderColor: '#ccc',
    borderWidth: 1
  },
  button: {
    backgroundColor: '#F4C542',
    paddingVertical: 12,
    paddingHorizontal: 25,
    borderRadius: 10,
    marginTop: 10
  },
  secondaryButton: {
    backgroundColor: '#A2D729'
  },
  buttonText: {
    color: '#fff',
    fontSize: 16,
    fontWeight: '600'
  }
});
`;

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

const moveDirectories = async (userInput) => {
  try {
    if (userInput === "y") {
      await fs.promises.mkdir(exampleDirPath, { recursive: true });
      console.log(`📁 /${exampleDir} directory created.`);
    }

    // Move ou deleta as pastas antigas
    for (const dir of oldDirs) {
      const oldDirPath = path.join(root, dir);
      if (fs.existsSync(oldDirPath)) {
        if (userInput === "y") {
          const newDirPath = path.join(exampleDirPath, dir);
          await fs.promises.rename(oldDirPath, newDirPath);
          console.log(`➡️ /${dir} moved to /${exampleDir}/${dir}.`);
        } else {
          await fs.promises.rm(oldDirPath, { recursive: true, force: true });
          console.log(`❌ /${dir} deleted.`);
        }
      }
    }

    // Cria as novas pastas em src/
    for (const dir of newDirs) {
      const dirPath = path.join(root, dir);
      await fs.promises.mkdir(dirPath, { recursive: true });
      console.log(`📁 /${dir} directory created.`);
    }

    // Cria src/app/index.tsx
    const indexPath = path.join(srcDir, "app", "index.tsx");
    await fs.promises.writeFile(indexPath, indexContent);
    console.log("📄 src/app/index.tsx created.");

    // Cria src/app/_layout.tsx
    const layoutPath = path.join(srcDir, "app", "_layout.tsx");
    await fs.promises.writeFile(layoutPath, layoutContent);
    console.log("📄 src/app/_layout.tsx created.");

    // Cria src/components/SolarFacilApp.tsx
    const componentPath = path.join(srcDir, "components", "SolarFacilApp.tsx");
    await fs.promises.writeFile(componentPath, componentContent);
    console.log("📄 src/components/SolarFacilApp.tsx created.");

    console.log("\n✅ Project reset complete. Next steps:");
    console.log(
      `1. Run \`npx expo start\` to start a development server.\n2. Edit src/app/index.tsx to edit the main screen.${
        userInput === "y"
          ? `\n3. Delete the /${exampleDir} directory when you're done referencing it.`
          : ""
      }`,
    );
  } catch (error) {
    console.error(`❌ Error during script execution: ${error.message}`);
  }
};

rl.question(
  "Do you want to move existing files to /app-example instead of deleting them? (Y/n): ",
  (answer) => {
    const userInput = answer.trim().toLowerCase() || "y";
    if (userInput === "y" || userInput === "n") {
      moveDirectories(userInput).finally(() => rl.close());
    } else {
      console.log("❌ Invalid input. Please enter 'Y' or 'N'.");
      rl.close();
    }
  },
);
