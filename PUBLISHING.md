# Publishing Guide

## How Publishing Works

This project automatically publishes to **npm** when you create a GitHub release. Once published to npm, **jsDelivr** automatically mirrors the package within minutes.

### Distribution Chain
```
GitHub Release → npm → jsDelivr CDN
```

## Quick Start: Publishing a New Version

1. **Make your changes** and ensure they pass validation
2. **Create a GitHub release**:
   - Go to https://github.com/afaf-tech/sholawat-json/releases/new
   - Choose a version tag (e.g., `v0.3.2` or `0.3.2`)
   - Write release notes
   - Click "Publish release"
3. **Automated workflow runs**:
   - ✅ Validates all JSON files
   - ✅ Updates `package.json` version
   - ✅ Publishes to npm
4. **Wait ~5 minutes** for jsDelivr to pick up the new version

## Prerequisites (One-Time Setup)

### 1. Create npm Account
If you don't have one: https://www.npmjs.com/signup

### 2. Get npm Access Token
1. Log in to npm: https://www.npmjs.com/
2. Click your profile → "Access Tokens"
3. Click "Generate New Token" → "Classic Token"
4. Select **Automation** type
5. Copy the token (starts with `npm_...`)

### 3. Add Token to GitHub
1. Go to: https://github.com/afaf-tech/sholawat-json/settings/secrets/actions
2. Click "New repository secret"
3. Name: `NPM_TOKEN`
4. Value: Paste your npm token
5. Click "Add secret"

## Version Numbering

Follow [Semantic Versioning](https://semver.org/):

- **Patch** (e.g., `0.3.1` → `0.3.2`): Bug fixes, typo corrections
- **Minor** (e.g., `0.3.1` → `0.4.0`): New content, new translations, backwards-compatible
- **Major** (e.g., `0.3.1` → `1.0.0`): Breaking changes to JSON structure/schema

## Creating a Release

### Via GitHub Web Interface

1. Go to: https://github.com/afaf-tech/sholawat-json/releases/new
2. **Tag version**: Enter version (e.g., `v0.3.2` or `0.3.2`)
3. **Target**: Keep as `master`
4. **Release title**: Same as tag (e.g., `v0.3.2`)
5. **Description**: Summarize changes:
   ```markdown
   ## What's New
   - Added Indonesian translation for Burdah Fasl 10
   - Fixed typo in Simtudduror Fasl 5
   
   ## Files Changed
   - sholawat/burdah/nu_online/fasl/10.json
   - sholawat/simtudduror/nu_online/fasl/5.json
   ```
6. Click **"Publish release"**

### Via Command Line (Alternative)

```bash
# Tag the current commit
git tag v0.3.2

# Push the tag
git push origin v0.3.2

# Then create release on GitHub using the tag
```

## What Happens During Publish

1. **GitHub Actions triggers** (`.github/workflows/publish.yml`)
2. **Validation runs** using Go validator
   - If validation fails, publish stops
3. **Version updated** in `package.json` based on release tag
4. **Published to npm** with provenance attestation
5. **jsDelivr mirrors** automatically within 5-10 minutes

## Verifying Publication

### Check npm
```bash
npm view sholawat-json version
```

Or visit: https://www.npmjs.com/package/sholawat-json

### Check jsDelivr
- Latest version: https://cdn.jsdelivr.net/npm/sholawat-json@latest/sholawat/sholawat.json
- Specific version: https://cdn.jsdelivr.net/npm/sholawat-json@0.3.2/sholawat/sholawat.json
- Purge cache: https://www.jsdelivr.com/tools/purge

## Manual Publishing (Fallback)

If automated workflow fails:

```bash
# 1. Validate locally
cd validator
go run main.go

# 2. Update version in package.json
npm version patch  # or minor, major

# 3. Login to npm (one time)
npm login

# 4. Publish
npm publish --access public
```

## Continuous Integration

### Validation Workflow (`.github/workflows/validate.yml`)

Runs on every push/PR to master that modifies JSON files:
- ✅ Validates all JSON against schemas
- ✅ Prevents invalid data from being merged
- ✅ Ensures data integrity before release

### When Validation Fails

1. Check the Actions tab: https://github.com/afaf-tech/sholawat-json/actions
2. Click the failed workflow run
3. Read the error messages from validator
4. Fix the JSON files
5. Commit and push fixes

## Best Practices

1. **Always validate locally** before pushing:
   ```bash
   cd validator && go run main.go
   ```

2. **Test changes** in a branch before merging to master

3. **Write clear release notes** describing what changed

4. **Use semantic versioning** consistently

5. **Don't delete releases** - users may depend on specific versions

6. **Update sholawat.json** when adding new files or sources

## Troubleshooting

### "NPM_TOKEN not found"
- Ensure the secret is added in repository settings
- Name must be exactly `NPM_TOKEN`

### "Validation failed"
- Run `cd validator && go run main.go` locally
- Fix reported schema violations
- Commit fixes and re-release

### "Permission denied" when publishing
- Verify npm token has correct permissions
- Ensure you're a maintainer on npm package

### jsDelivr shows old version
- Wait 5-10 minutes for CDN propagation
- Force purge: https://www.jsdelivr.com/tools/purge
- Use specific version URL instead of `@latest`

## Resources

- npm package: https://www.npmjs.com/package/sholawat-json
- jsDelivr CDN: https://www.jsdelivr.com/package/npm/sholawat-json
- GitHub Actions: https://github.com/afaf-tech/sholawat-json/actions
- Semantic Versioning: https://semver.org/

## Questions?

Open an issue: https://github.com/afaf-tech/sholawat-json/issues
