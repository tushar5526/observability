// This file was generated by https://github.com/grafana/dashboard-spec

{
  new(
    collapse=true,
    collapsed=true,
    datasource=null,
    repeat=null,
    repeatIteration=null,
    showTitle=true,
    title=null,
    titleSize='h6',
  ):: {
    [if collapse != null then 'collapse']: collapse,
    [if collapsed != null then 'collapsed']: collapsed,
    [if datasource != null then 'datasource']: datasource,
    [if repeat != null then 'repeat']: repeat,
    [if repeatIteration != null then 'repeatIteration']: repeatIteration,
    [if showTitle != null then 'showTitle']: showTitle,
    [if title != null then 'title']: title,
    [if titleSize != null then 'titleSize']: titleSize,
    type: 'row',

    setGridPos(
      h=8,
      w=12,
      x=null,
      y=null,
    ):: self {}
        + { gridPos+: { [if h != null then 'h']: h } }
        + { gridPos+: { [if w != null then 'w']: w } }
        + { gridPos+: { [if x != null then 'x']: x } }
        + { gridPos+: { [if y != null then 'y']: y } }
    ,


    addPanel(
      panel
    ):: self {}
        + { panels+: [
          panel,
        ] },

  },
}