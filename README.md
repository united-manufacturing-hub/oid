# OID

This project contains our OID definitions. We use [IANA's prefix and our PEN](https://www.iana.org/assignments/enterprise-numbers/enterprise-numbers) for our OIDs.

## How to use

If you want to contribute to this repository make sure to set up our githooks:
```bash
git config core.hooksPath .githooks
```

Make sure to always git tag your code changes. We use [semantic versioning](https://semver.org/).

## Contributing

```
The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL
      NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED",  "MAY", and
      "OPTIONAL" in this document are to be interpreted as described in
RFC 2119.
```

### Definitions

- Root OID:
   - This is the IANA prefix + UMH root OID
   - 1.3.6.1.4.1.59193
   - This OID **[must not](https://datatracker.ietf.org/doc/html/rfc2119#section-2)** be used on its own
- Intermediate OID
   - An OID, exactly one level below the Root OID
- Sub OID:
   - An OID, any level below an Intermediate OID

### Rules

1. Intermediate OID **[must](https://datatracker.ietf.org/doc/html/rfc2119#section-1)** be increments of 100, for easier reading
2. There **[must](https://datatracker.ietf.org/doc/html/rfc2119#section-1)** only be one Intermediate OID per application.
3. If there is a need for multiple Intermediate OID per application, an application **[may](https://datatracker.ietf.org/doc/html/rfc2119#section-5)** have one or multiple sub-OID
4. Intermediate OID **[must not](https://datatracker.ietf.org/doc/html/rfc2119#section-2)** be deleted
5. Sub OID **[should not](https://datatracker.ietf.org/doc/html/rfc2119#section-4)** be deleted or modified, except for mayor version upgrades, of the application.
    1. Any such deletion or modification **[must](https://datatracker.ietf.org/doc/html/rfc2119#section-1)** are shown within the commit history
6. Registrations and marking of modifications & deletions to the Development Intermediate OID Sub OID are **[optional](https://datatracker.ietf.org/doc/html/rfc2119#section-5)**
7. All OIDs **[must](https://datatracker.ietf.org/doc/html/rfc2119#section-1)**, except for the development OID, be implemented in this repository.