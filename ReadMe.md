# RPKI Validator & ASN Prefix API

This is a simple api for validating RPKI records for an ASN's entire list of announced prefixes.

## API Routes

- /api/rpki?q={as number}

  - @Returns [{ prefix, origin, rpki_state, description }]

- /api/prefixes?q={as number}

  - @Returns [{ prefix, origin_asn }]
