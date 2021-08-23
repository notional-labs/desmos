/**
 * Creating a sidebar enables you to:
 - create an ordered group of docs
 - render a sidebar for each doc of that group
 - provide next/previous navigation

 The sidebars can be generated from the filesystem, or explicitly defined here.

 Create as many sidebars as you want.
 */

module.exports = {
  docs: [
      'intro',
      {
        type: 'category',
        label: 'Developers',
        collapsed: false,
        items: [
            'developers/overview',
            'developers/types',
            'developers/perform-transactions',
            'developers/query-data',
            'developers/observe-data',
            'developers/developer-faq',
        ]
      },
      {
          type: 'category',
          label: 'Running a Fullnode',
          collapsed: false,
          items: [
              'running-a-full-node/overview',
              'running-a-full-node/setup',
              'running-a-full-node/update',
              'running-a-full-node/rocksdb-installation',
              'running-a-full-node/cosmovisor'
          ]
      },
      {
          type: 'category',
          label: 'Validators',
          collapsed: false,
          items: [
              'validators/overview',
              'validators/security',
              'validators/setup',
              'validators/halting',
              'validators/migrating',
              'validators/common-problems',
              'validators/validator-faq'
          ]
      },
      {
          type: 'category',
          label: 'Testnets',
          collapsed: false,
          items: [
              'testnets/overview',
              'testnets/create-local',
              'testnets/join-public'
          ]
      }
      /*{
        type: 'autogenerated',
        dirName: '.'
      }*/
  ],
};
