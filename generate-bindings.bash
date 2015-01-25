#!/bin/bash

mkdir -p android/sdk

( cd api-doc-parser && go build )
( cd api-code-generation && go build )

(
    api-doc-parser/api-doc-parser -g '^View$' view ;
    api-doc-parser/api-doc-parser -g 'View$' widget
) | api-code-generation/api-code-generation android/sdk -b View -r ViewTypeConstructors

(
    api-doc-parser/api-doc-parser -g '^Button$' widget ;
    api-doc-parser/api-doc-parser -g '^EditText$' widget
) | api-code-generation/api-code-generation android/sdk -b TextView -r ViewTypeConstructors

(
    api-doc-parser/api-doc-parser -g '^ToggleButton$' widget ;
) | api-code-generation/api-code-generation android/sdk -b Button -r ViewTypeConstructors
