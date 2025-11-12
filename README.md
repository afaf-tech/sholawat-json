# Sholawat JSON

All praise is due to Allah SWT for His countless blessings. May blessings and peace always be upon our beloved Prophet Muhammad SAW.

**sholawat-json** is a collection of sholawat in a simple JSON format, aimed at developers who wish to integrate them into websites or applications. Each JSON file represents a different sholawat and follows a predefined schema, ensuring ease of access and consistent structure. To maintain accuracy and structure, we implement `JSON schema` validation, making it easier to integrate the data confidently. This project is designed for lightweight applications, with a focus on speed and efficiency.

We strive for accuracy, but if you encounter any errors, please reach out. Your feedback is invaluable for improvement.

## 🚀 Quick Start

### Option 1: CDN (Recommended for Web Projects)

No installation needed! Just fetch directly from jsDelivr CDN:

```javascript
// Fetch list of all available sholawat
fetch('https://cdn.jsdelivr.net/npm/sholawat-json@latest/sholawat/sholawat.json')
  .then(response => response.json())
  .then(data => console.log(data));
```

### Option 2: npm Package

```bash
npm install sholawat-json
```

Then import in your project:

```javascript
import sholawatList from 'sholawat-json/sholawat/sholawat.json';
import burdahFasl1 from 'sholawat-json/sholawat/burdah/nu_online/fasl/1.json';
```

---

## 📖 Usage Guide

### Step 1: Get Available Sholawat

First, fetch the master list to see all available sholawat, sources, and file paths:

**Endpoint:**
```
https://cdn.jsdelivr.net/npm/sholawat-json@latest/sholawat/sholawat.json
```

**Example Response:**
```json
[
  {
    "name": "burdah",
    "sources": [
      {
        "source_name": "nu online",
        "description": "burdah-fasl",
        "path_files": "sholawat/burdah/nu_online/fasl/",
        "files": ["1.json", "2.json", "3.json", ...],
        "schema": "burdah_fasl_v1.0.schema.json"
      }
    ]
  }
]
```

### Step 2: Fetch Specific Sholawat Content

#### For Fasl-based Sholawat (Burdah, Diba, Simtudduror)

**URL Pattern:**
```
https://cdn.jsdelivr.net/npm/sholawat-json@latest/sholawat/{name}/{source}/fasl/{number}.json
```

**Examples:**

```javascript
// Fetch Burdah Fasl 1 from NU Online
const burdahUrl = 'https://cdn.jsdelivr.net/npm/sholawat-json@latest/sholawat/burdah/nu_online/fasl/1.json';

// Fetch Diba Fasl 10 from NU Online
const dibaUrl = 'https://cdn.jsdelivr.net/npm/sholawat-json@latest/sholawat/diba/nu_online/fasl/10.json';

// Fetch Simtudduror Fasl 5
const simtuddurorUrl = 'https://cdn.jsdelivr.net/npm/sholawat-json@latest/sholawat/simtudduror/nu_online/fasl/5.json';
```

#### For Single Sholawat (Tunggal & Suluk)

**URL Pattern:**
```
https://cdn.jsdelivr.net/npm/sholawat-json@latest/sholawat/{type}/{filename}.json
```

**Examples:**

```javascript
// Fetch a tunggal sholawat
const tunggalUrl = 'https://cdn.jsdelivr.net/npm/sholawat-json@latest/sholawat/tunggal/assalamu-alaika-zainal-anbiya.json';

// Fetch a suluk sholawat
const sulukUrl = 'https://cdn.jsdelivr.net/npm/sholawat-json@latest/sholawat/suluk/hubbun-nabi.json';
```

---

## 💻 Code Examples

### Vanilla JavaScript (Fetch API)

```javascript
async function getSholawat() {
  try {
    const response = await fetch(
      'https://cdn.jsdelivr.net/npm/sholawat-json@latest/sholawat/burdah/nu_online/fasl/1.json'
    );
    const data = await response.json();
    
    console.log('Sholawat Name:', data.name);
    console.log('Arabic Text:', data.text);
    console.log('Translation:', data.translations.id.name);
  } catch (error) {
    console.error('Error fetching sholawat:', error);
  }
}

getSholawat();
```

### React Example

```jsx
import React, { useState, useEffect } from 'react';

function SholawatDisplay() {
  const [sholawat, setSholawat] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch('https://cdn.jsdelivr.net/npm/sholawat-json@latest/sholawat/diba/nu_online/fasl/1.json')
      .then(res => res.json())
      .then(data => {
        setSholawat(data);
        setLoading(false);
      })
      .catch(err => console.error(err));
  }, []);

  if (loading) return <div>Loading...</div>;

  return (
    <div>
      <h2>{sholawat.name}</h2>
      <div>
        {Object.entries(sholawat.text).map(([key, verse]) => (
          <div key={key}>
            <p className="arabic">{verse.arabic}</p>
            <p className="latin">{verse.latin}</p>
          </div>
        ))}
      </div>
    </div>
  );
}
```

### Vue.js Example

```vue
<template>
  <div v-if="sholawat">
    <h2>{{ sholawat.name }}</h2>
    <div v-for="(verse, key) in sholawat.text" :key="key">
      <p class="arabic">{{ verse.arabic }}</p>
      <p class="latin">{{ verse.latin }}</p>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      sholawat: null
    };
  },
  mounted() {
    fetch('https://cdn.jsdelivr.net/npm/sholawat-json@latest/sholawat/burdah/nu_online/fasl/1.json')
      .then(res => res.json())
      .then(data => this.sholawat = data);
  }
};
</script>
```

### Axios Example

```javascript
import axios from 'axios';

const API_BASE = 'https://cdn.jsdelivr.net/npm/sholawat-json@latest';

// Get all sholawat list
const getAllSholawat = () => 
  axios.get(`${API_BASE}/sholawat/sholawat.json`);

// Get specific fasl
const getBurdahFasl = (number) => 
  axios.get(`${API_BASE}/sholawat/burdah/nu_online/fasl/${number}.json`);

// Usage
getAllSholawat()
  .then(response => console.log(response.data))
  .catch(error => console.error(error));

getBurdahFasl(1)
  .then(response => console.log(response.data))
  .catch(error => console.error(error));
```

---

## 📦 Data Structure

### Fasl-based Sholawat (Burdah, Diba, Simtudduror)

```json
{
  "number": 1,
  "source": "nu_online",
  "name": "بَانَتْ سُعَادُ",
  "text": {
    "1": {
      "arabic": "بَانَتْ سُعَادُ فَقَلْبِيْ الْيَوْمَ مَتْبُوْلُ",
      "latin": "Bānat Su'ādu faqalbil yauma matbūlu"
    }
  },
  "translations": {
    "id": {
      "name": "Suad Telah Pergi",
      "text": {
        "1": "Suad telah pergi, maka hatiku hari ini terguncang"
      }
    }
  },
  "last_updated": "2024-06-15"
}
```

### Single Sholawat (Tunggal & Suluk)

```json
{
  "source": "general",
  "name": "يَا سَيِّدِيْ يَا رَسُوْلَ اللهِ",
  "latin": "Ya Sayyidi Ya Rasulallah",
  "text": {
    "1": {
      "arabic": "يَا سَيِّدِيْ يَا رَسُوْلَ اللهِ",
      "latin": "Yā sayyidī yā rasūlallāh"
    }
  },
  "translations": {
    "id": {
      "name": "Ya Tuanku Ya Rasulullah",
      "translator": "Tim NU Online",
      "text": {
        "1": "Wahai tuanku wahai Rasulullah"
      }
    }
  },
  "last_updated": "2024-07-04"
}
```

---

## 🔍 Schema Validation

Each sholawat type follows a JSON Schema for data consistency. You can access schemas at:

```
https://cdn.jsdelivr.net/npm/sholawat-json@latest/schemas/{schema_name}.schema.json
```

**Available Schemas:**
- `burdah_fasl_v1.0.schema.json` - Burdah chapters
- `diba_fasl_v1.0.schema.json` - Diba chapters  
- `simtudduror_fasl_v1.0.schema.json` - Simtudduror chapters
- `sholawat_tunggal_v1.0.schema.json` - Single sholawat
- `suluk_v1.0.schema.json` - Suluk sholawat

**Example:**
```javascript
const schemaUrl = 'https://cdn.jsdelivr.net/npm/sholawat-json@latest/schemas/burdah_fasl_v1.0.schema.json';
```

---

## 🗂️ Browse Files

Explore all available sholawat files in the repository:
- [Browse `/sholawat` directory](https://github.com/afaf-tech/sholawat-json/tree/master/sholawat)
- [View on jsDelivr CDN](https://www.jsdelivr.com/package/npm/sholawat-json)

## Data Sources

The sholawat data in this project has been collected from various trusted sources, including reputable Islamic literature and applications. Each data entry has gone through a multi-stage validation process to ensure accuracy and authenticity.

## Contributions

We warmly welcome contributions from anyone who has ideas or suggestions for the development of this project. If you would like to contribute, please submit a Pull Request (PR) through GitHub.

## License

This project is licensed under the MIT License. For more details, please refer to the LICENSE file included in the repository.

## Author
This project was developed by afaf-tech. If you encounter any issues, please report them by creating a GitHub issue in the repository.